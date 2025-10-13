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

// TODO : Objects to be present in the grid for testing
// - DDNS Principal Cluster Group - example-ddns-principal-group, updated-ddns-principal-group
// - Grid Members - infoblox.localdomain1, infoblox.localdomain2
// - NS Group - example-ns-group, updated-example-ns-group
// - Record Name Policy - example-policy, example-policy-update
// - MS Server - 10.120.23.22, 10.120.23.23
// - Shared Record Group - example-shared-record-group, updated-example-shared-record-groups
// - GSS TSIG Key has to be configured in the grid

var readableAttributesForZoneAuth = "address,allow_active_dir,allow_fixed_rrset_order,allow_gss_tsig_for_underscore_zone,allow_gss_tsig_zone_updates,allow_query,allow_transfer,allow_update,allow_update_forwarding,aws_rte53_zone_info,cloud_info,comment,copy_xfer_to_notify,create_underscore_zones,ddns_force_creation_timestamp_update,ddns_principal_group,ddns_principal_tracking,ddns_restrict_patterns,ddns_restrict_patterns_list,ddns_restrict_protected,ddns_restrict_secure,ddns_restrict_static,disable,disable_forwarding,display_domain,dns_fqdn,dns_integrity_enable,dns_integrity_frequency,dns_integrity_member,dns_integrity_verbose_logging,dns_soa_email,dnssec_key_params,dnssec_keys,dnssec_ksk_rollover_date,dnssec_zsk_rollover_date,effective_check_names_policy,effective_record_name_policy,extattrs,external_primaries,external_secondaries,fqdn,grid_primary,grid_primary_shared_with_ms_parent_delegation,grid_secondaries,is_dnssec_enabled,is_dnssec_signed,is_multimaster,last_queried,last_queried_acl,locked,locked_by,mask_prefix,member_soa_mnames,member_soa_serials,ms_ad_integrated,ms_allow_transfer,ms_allow_transfer_mode,ms_dc_ns_record_creation,ms_ddns_mode,ms_managed,ms_primaries,ms_read_only,ms_secondaries,ms_sync_disabled,ms_sync_master_name,network_associations,network_view,notify_delay,ns_group,parent,prefix,primary_type,record_name_policy,records_monitored,rr_not_queried_enabled_time,scavenging_settings,soa_default_ttl,soa_email,soa_expire,soa_negative_ttl,soa_refresh,soa_retry,soa_serial_number,srgs,update_forwarding,use_allow_active_dir,use_allow_query,use_allow_transfer,use_allow_update,use_allow_update_forwarding,use_check_names_policy,use_copy_xfer_to_notify,use_ddns_force_creation_timestamp_update,use_ddns_patterns_restriction,use_ddns_principal_security,use_ddns_restrict_protected,use_ddns_restrict_static,use_dnssec_key_params,use_external_primary,use_grid_zone_timer,use_import_from,use_notify_delay,use_record_name_policy,use_scavenging_settings,use_soa_email,using_srg_associations,view,zone_format,zone_not_queried_enabled_time"

func TestAccZoneAuthResource_basic(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("auth-zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthBasicConfig(zoneFqdn, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fqdn", zoneFqdn),
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
					// Test fields with default values
					resource.TestCheckResourceAttr(resourceName, "allow_fixed_rrset_order", "false"),
					resource.TestCheckResourceAttr(resourceName, "allow_gss_tsig_for_underscore_zone", "false"),
					resource.TestCheckResourceAttr(resourceName, "allow_gss_tsig_zone_updates", "false"),
					resource.TestCheckResourceAttr(resourceName, "copy_xfer_to_notify", "false"),
					resource.TestCheckResourceAttr(resourceName, "create_ptr_for_bulk_hosts", "false"),
					resource.TestCheckResourceAttr(resourceName, "create_ptr_for_hosts", "false"),
					resource.TestCheckResourceAttr(resourceName, "create_underscore_zones", "false"),
					resource.TestCheckResourceAttr(resourceName, "ddns_force_creation_timestamp_update", "false"),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal_tracking", "false"),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_patterns", "false"),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_protected", "false"),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_secure", "false"),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_static", "false"),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
					resource.TestCheckResourceAttr(resourceName, "disable_forwarding", "false"),
					resource.TestCheckResourceAttr(resourceName, "dns_integrity_enable", "false"),
					resource.TestCheckResourceAttr(resourceName, "dns_integrity_frequency", "3600"),
					resource.TestCheckResourceAttr(resourceName, "dns_integrity_verbose_logging", "false"),
					resource.TestCheckResourceAttr(resourceName, "effective_check_names_policy", "WARN"),
					resource.TestCheckResourceAttr(resourceName, "locked", "false"),
					resource.TestCheckResourceAttr(resourceName, "ms_ad_integrated", "false"),
					resource.TestCheckResourceAttr(resourceName, "ms_allow_transfer_mode", "NONE"),
					resource.TestCheckResourceAttr(resourceName, "ms_ddns_mode", "NONE"),
					resource.TestCheckResourceAttr(resourceName, "ms_sync_disabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "notify_delay", "5"),
					resource.TestCheckResourceAttr(resourceName, "use_allow_active_dir", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_allow_query", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_allow_transfer", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_allow_update", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_check_names_policy", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_restrict_static", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_dnssec_key_params", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_external_primary", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_import_from", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_notify_delay", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_record_name_policy", "false"),
					resource.TestCheckResourceAttr(resourceName, "use_scavenging_settings", "false"),
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
					resource.TestCheckResourceAttr(resourceName, "zone_format", "FORWARD"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_disappears(t *testing.T) {
	resourceName := "nios_dns_zone_auth.test"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("auth-zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckZoneAuthDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccZoneAuthBasicConfig(zoneFqdn, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					testAccCheckZoneAuthDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccZoneAuthResource_AllowActiveDir(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_allow_active_dir"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"
	allowActiveDir := []map[string]any{
		{
			"address":    "10.0.0.1",
			"permission": "ALLOW",
		},
	}

	updatedAllowActiveDir := []map[string]any{
		{
			"address":    "10.0.0.2",
			"permission": "ALLOW",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthAllowActiveDir(zoneFqdn, "default", allowActiveDir, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_active_dir.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "allow_active_dir.0.address", "10.0.0.1"),
					resource.TestCheckResourceAttr(resourceName, "allow_active_dir.0.permission", "ALLOW"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthAllowActiveDir(zoneFqdn, "default", updatedAllowActiveDir, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_active_dir.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "allow_active_dir.0.address", "10.0.0.2"),
					resource.TestCheckResourceAttr(resourceName, "allow_active_dir.0.permission", "ALLOW"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_AllowFixedRrsetOrder(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_allow_fixed_rrset_order"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthAllowFixedRrsetOrder(zoneFqdn, "default", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_fixed_rrset_order", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthAllowFixedRrsetOrder(zoneFqdn, "default", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_fixed_rrset_order", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_AllowGssTsigForUnderscoreZone(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_allow_gss_tsig_for_underscore_zone"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthAllowGssTsigForUnderscoreZone(zoneFqdn, "default", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_gss_tsig_for_underscore_zone", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthAllowGssTsigForUnderscoreZone(zoneFqdn, "default", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_gss_tsig_for_underscore_zone", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_AllowGssTsigZoneUpdates(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_allow_gss_tsig_zone_updates"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthAllowGssTsigZoneUpdates(zoneFqdn, "default", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_gss_tsig_zone_updates", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthAllowGssTsigZoneUpdates(zoneFqdn, "default", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_gss_tsig_zone_updates", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_AllowQuery(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_allow_query"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"
	allowQuery := []map[string]any{
		{
			"struct":     "addressac",
			"address":    "10.0.0.0",
			"permission": "ALLOW",
		},
	}
	updatedAllowQuery := []map[string]any{
		{
			"struct":     "addressac",
			"address":    "192.168.0.0",
			"permission": "ALLOW",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthAllowQuery(zoneFqdn, "default", allowQuery, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_query.0.address", "10.0.0.0"),
					resource.TestCheckResourceAttr(resourceName, "allow_query.0.permission", "ALLOW"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthAllowQuery(zoneFqdn, "default", updatedAllowQuery, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_query.0.address", "192.168.0.0"),
					resource.TestCheckResourceAttr(resourceName, "allow_query.0.permission", "ALLOW"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_AllowTransfer(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_allow_transfer"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"
	allowTransfer := []map[string]any{
		{
			"struct":     "addressac",
			"address":    "10.0.0.0",
			"permission": "ALLOW",
		},
	}
	updatedAllowTransfer := []map[string]any{
		{
			"struct":     "addressac",
			"address":    "192.168.0.0",
			"permission": "ALLOW",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthAllowTransfer(zoneFqdn, "default", allowTransfer, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_transfer.0.address", "10.0.0.0"),
					resource.TestCheckResourceAttr(resourceName, "allow_transfer.0.permission", "ALLOW"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthAllowTransfer(zoneFqdn, "default", updatedAllowTransfer, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_transfer.0.address", "192.168.0.0"),
					resource.TestCheckResourceAttr(resourceName, "allow_transfer.0.permission", "ALLOW"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_AllowUpdate(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_allow_update"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"
	allowUpdate := []map[string]any{
		{
			"struct":     "addressac",
			"address":    "10.0.0.0",
			"permission": "ALLOW",
		},
	}
	updatedAllowUpdate := []map[string]any{
		{
			"struct":     "addressac",
			"address":    "192.168.0.0",
			"permission": "ALLOW",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthAllowUpdate(zoneFqdn, "default", allowUpdate, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_update.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "allow_update.0.struct", "addressac"),
					resource.TestCheckResourceAttr(resourceName, "allow_update.0.address", "10.0.0.0"),
					resource.TestCheckResourceAttr(resourceName, "allow_update.0.permission", "ALLOW"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthAllowUpdate(zoneFqdn, "default", updatedAllowUpdate, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_update.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "allow_update.0.struct", "addressac"),
					resource.TestCheckResourceAttr(resourceName, "allow_update.0.address", "192.168.0.0"),
					resource.TestCheckResourceAttr(resourceName, "allow_update.0.permission", "ALLOW"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_AllowUpdateForwarding(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_allow_update_forwarding"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthAllowUpdateForwarding(zoneFqdn, "default", true, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_update_forwarding", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthAllowUpdateForwarding(zoneFqdn, "default", false, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_update_forwarding", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_Comment(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_comment"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthComment(zoneFqdn, "default", "initial comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "initial comment"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthComment(zoneFqdn, "default", "updated comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_CopyXferToNotify(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_copy_xfer_to_notify"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthCopyXferToNotify(zoneFqdn, "default", true, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "copy_xfer_to_notify", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthCopyXferToNotify(zoneFqdn, "default", false, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "copy_xfer_to_notify", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_CreateUnderscoreZones(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_create_underscore_zones"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthCreateUnderscoreZones(zoneFqdn, "default", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "create_underscore_zones", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthCreateUnderscoreZones(zoneFqdn, "default", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "create_underscore_zones", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_DdnsForceCreationTimestampUpdate(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_ddns_force_creation_timestamp_update"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthDdnsForceCreationTimestampUpdate(zoneFqdn, "default", true, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_force_creation_timestamp_update", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthDdnsForceCreationTimestampUpdate(zoneFqdn, "default", false, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_force_creation_timestamp_update", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_DdnsPrincipalGroup(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_ddns_principal_group"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthDdnsPrincipalGroup(zoneFqdn, "default", "example-ddns-principal-group", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal_group", "example-ddns-principal-group"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthDdnsPrincipalGroup(zoneFqdn, "default", "updated-ddns-principal-group", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal_group", "updated-ddns-principal-group"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_DdnsPrincipalTracking(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_ddns_principal_tracking"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthDdnsPrincipalTracking(zoneFqdn, "default", true, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal_tracking", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthDdnsPrincipalTracking(zoneFqdn, "default", false, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal_tracking", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_DdnsRestrictPatterns(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_ddns_restrict_patterns"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthDdnsRestrictPatterns(zoneFqdn, "default", true, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_patterns", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthDdnsRestrictPatterns(zoneFqdn, "default", false, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_patterns", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_DdnsRestrictPatternsList(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_ddns_restrict_patterns_list"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"
	patternList := []string{"pattern1", "pattern2"}
	updatedPatternList := []string{"pattern3", "pattern4"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthDdnsRestrictPatternsList(zoneFqdn, "default", patternList, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_patterns_list.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_patterns_list.0", "pattern1"),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_patterns_list.1", "pattern2"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthDdnsRestrictPatternsList(zoneFqdn, "default", updatedPatternList, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_patterns_list.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_patterns_list.0", "pattern3"),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_patterns_list.1", "pattern4"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_DdnsRestrictProtected(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_ddns_restrict_protected"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthDdnsRestrictProtected(zoneFqdn, "default", true, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_protected", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthDdnsRestrictProtected(zoneFqdn, "default", false, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_protected", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_DdnsRestrictSecure(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_ddns_restrict_secure"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthDdnsRestrictSecure(zoneFqdn, "default", true, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_secure", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthDdnsRestrictSecure(zoneFqdn, "default", false, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_secure", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_DdnsRestrictStatic(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_ddns_restrict_static"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthDdnsRestrictStatic(zoneFqdn, "default", true, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_static", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthDdnsRestrictStatic(zoneFqdn, "default", false, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_static", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_Disable(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_disable"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthDisable(zoneFqdn, "default", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthDisable(zoneFqdn, "default", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_DisableForwarding(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_disable_forwarding"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthDisableForwarding(zoneFqdn, "default", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable_forwarding", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthDisableForwarding(zoneFqdn, "default", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable_forwarding", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_DnsIntegrityEnable(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_dns_integrity_enable"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"
	gridPrimary := []map[string]any{
		{
			"name": "infoblox.localdomain2",
		},
	}
	dnsIntegrityMember := "infoblox.localdomain2"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthDnsIntegrityEnable(zoneFqdn, "default", true, gridPrimary, dnsIntegrityMember),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns_integrity_enable", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthDnsIntegrityEnable(zoneFqdn, "default", false, gridPrimary, dnsIntegrityMember),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns_integrity_enable", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_DnsIntegrityFrequency(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_dns_integrity_frequency"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthDnsIntegrityFrequency(zoneFqdn, "default", 1000),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns_integrity_frequency", "1000"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthDnsIntegrityFrequency(zoneFqdn, "default", 2000),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns_integrity_frequency", "2000"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_DnsIntegrityMember(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_dns_integrity_member"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"
	dnsIntegrityMember1 := "infoblox.localdomain2"
	dnsIntegrityMember2 := "infoblox.localdomain1"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthDnsIntegrityMember(zoneFqdn, "default", dnsIntegrityMember1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns_integrity_member", dnsIntegrityMember1),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthDnsIntegrityMember(zoneFqdn, "default", dnsIntegrityMember2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns_integrity_member", dnsIntegrityMember2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_DnsIntegrityVerboseLogging(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_dns_integrity_verbose_logging"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthDnsIntegrityVerboseLogging(zoneFqdn, "default", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns_integrity_verbose_logging", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthDnsIntegrityVerboseLogging(zoneFqdn, "default", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns_integrity_verbose_logging", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_DnssecKeyParams(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_dnssec_key_params"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"
	kskAlgorithms := []map[string]any{
		{
			"algorithm": "RSASHA256",
			"size":      2048,
		},
	}
	zskAlgorithms := []map[string]any{
		{
			"algorithm": "RSASHA256",
			"size":      1024,
		},
	}
	updatedKskAlgorithms := []map[string]any{
		{
			"algorithm": "RSASHA512",
			"size":      4096,
		},
	}
	updatedZskAlgorithms := []map[string]any{
		{
			"algorithm": "RSASHA512",
			"size":      2048,
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthDnssecKeyParams(zoneFqdn, "default", kskAlgorithms, zskAlgorithms),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_dnssec_key_params", "true"),
					resource.TestCheckResourceAttr(resourceName, "dnssec_key_params.ksk_algorithms.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "dnssec_key_params.ksk_algorithms.0.algorithm", "RSASHA256"),
					resource.TestCheckResourceAttr(resourceName, "dnssec_key_params.ksk_algorithms.0.size", "2048"),
					resource.TestCheckResourceAttr(resourceName, "dnssec_key_params.zsk_algorithms.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "dnssec_key_params.zsk_algorithms.0.algorithm", "RSASHA256"),
					resource.TestCheckResourceAttr(resourceName, "dnssec_key_params.zsk_algorithms.0.size", "1024"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthDnssecKeyParams(zoneFqdn, "default", updatedKskAlgorithms, updatedZskAlgorithms),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_dnssec_key_params", "true"),
					resource.TestCheckResourceAttr(resourceName, "dnssec_key_params.ksk_algorithms.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "dnssec_key_params.ksk_algorithms.0.algorithm", "RSASHA512"),
					resource.TestCheckResourceAttr(resourceName, "dnssec_key_params.ksk_algorithms.0.size", "4096"),
					resource.TestCheckResourceAttr(resourceName, "dnssec_key_params.zsk_algorithms.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "dnssec_key_params.zsk_algorithms.0.algorithm", "RSASHA512"),
					resource.TestCheckResourceAttr(resourceName, "dnssec_key_params.zsk_algorithms.0.size", "2048"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_EffectiveCheckNamesPolicy(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_effective_check_names_policy"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthEffectiveCheckNamesPolicy(zoneFqdn, "default", "WARN"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "effective_check_names_policy", "WARN"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthEffectiveCheckNamesPolicy(zoneFqdn, "default", "FAIL"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "effective_check_names_policy", "FAIL"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_extattrs"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"
	extAttrValue1 := acctest.RandomName()
	extAttrValue2 := acctest.RandomName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthExtAttrs(zoneFqdn, "default", map[string]string{
					"Site": extAttrValue1,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue1),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthExtAttrs(zoneFqdn, "default", map[string]string{
					"Site": extAttrValue2,
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs.Site", extAttrValue2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_ExternalPrimaries(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_external_primaries"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"
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
			"name":    "example-server",
		},
	}
	msSecondaries := []map[string]any{
		{
			"address": "10.120.23.22",
			"ns_name": "example-server",
			"ns_ip":   "10.120.23.22",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthExternalPrimaries(zoneFqdn, "default", externalPrimaries, msSecondaries, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_primaries.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "external_primaries.0.address", "10.0.0.0"),
					resource.TestCheckResourceAttr(resourceName, "external_primaries.0.name", "example-server"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthExternalPrimaries(zoneFqdn, "default", updatedExternalPrimaries, msSecondaries, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_primaries.0.address", "10.0.0.2"),
					resource.TestCheckResourceAttr(resourceName, "external_primaries.0.name", "example-server"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_ExternalSecondaries(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_external_secondaries"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"
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
			"name": "infoblox.localdomain",
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
				Config: testAccZoneAuthExternalSecondaries(zoneFqdn, "default", externalSecondaries, gridPrimary),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_secondaries.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "external_secondaries.0.address", "10.0.0.0"),
					resource.TestCheckResourceAttr(resourceName, "external_secondaries.0.name", "example.com"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthExternalSecondaries(zoneFqdn, "default", updatedExternalSecondaries, gridPrimary),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_secondaries.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "external_secondaries.0.address", "10.0.0.2"),
					resource.TestCheckResourceAttr(resourceName, "external_secondaries.0.name", "updated-example.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_Fqdn(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_fqdn"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthFqdn(zoneFqdn, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fqdn", zoneFqdn),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthFqdn(zoneFqdn, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fqdn", zoneFqdn),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_GridPrimary(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_grid_primary"
	var v dns.ZoneAuth
	zoneAuth := acctest.RandomNameWithPrefix("zone") + ".com"
	gridPrimary := []map[string]any{
		{
			"name":    "infoblox.member",
			"stealth": false,
		},
	}
	gridPrimaryUpdated := []map[string]any{
		{
			"name":    "infoblox.localdomain",
			"stealth": true,
		},
		{
			"name":    "infoblox.member",
			"stealth": false,
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthGridPrimary(zoneAuth, "default", gridPrimary),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "grid_primary.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "grid_primary.0.name", "infoblox.member"),
					resource.TestCheckResourceAttr(resourceName, "grid_primary.0.stealth", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthGridPrimary(zoneAuth, "default", gridPrimaryUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "grid_primary.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "grid_primary.0.name", "infoblox.localdomain"),
					resource.TestCheckResourceAttr(resourceName, "grid_primary.0.stealth", "true"),
					resource.TestCheckResourceAttr(resourceName, "grid_primary.1.name", "infoblox.member"),
					resource.TestCheckResourceAttr(resourceName, "grid_primary.1.stealth", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_GridSecondaries(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_grid_secondaries"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"
	gridPrimary := []map[string]any{
		{
			"name": "infoblox.localdomain2",
		},
	}
	gridSecondary := []map[string]any{
		{
			"name": "infoblox.localdomain",
		},
	}
	updatedGridSecondary := []map[string]any{
		{
			"name": "infoblox.localdomain1",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthGridSecondaries(zoneFqdn, "default", gridPrimary, gridSecondary),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "grid_secondaries.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "grid_secondaries.0.name", "infoblox.localdomain"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthGridSecondaries(zoneFqdn, "default", gridPrimary, updatedGridSecondary),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "grid_secondaries.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "grid_secondaries.0.name", "infoblox.localdomain1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_LastQueriedAcl(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_last_queried_acl"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"
	lastQueriedACL := []map[string]any{
		{
			"address": "10.0.0.0",
		},
	}
	updatedLastQueriedACL := []map[string]any{
		{
			"address": "10.0.0.2",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthLastQueriedAcl(zoneFqdn, "default", lastQueriedACL, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "last_queried_acl.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "last_queried_acl.0.address", "10.0.0.0"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthLastQueriedAcl(zoneFqdn, "default", updatedLastQueriedACL, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "last_queried_acl.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "last_queried_acl.0.address", "10.0.0.2"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_Locked(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_locked"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthLocked(zoneFqdn, "default", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "locked", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthLocked(zoneFqdn, "default", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "locked", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_MemberSoaMnames(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_member_soa_mnames"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"
	gridPrimary := []map[string]any{
		{
			"name": "infoblox.localdomain",
		},
	}
	memberSoaMnames := []map[string]any{
		{
			"grid_primary": "infoblox.localdomain",
			"mname":        "infoblox.localdomain",
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
				Config: testAccZoneAuthMemberSoaMnames(zoneFqdn, "default", gridPrimary, memberSoaMnames),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "member_soa_mnames.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "member_soa_mnames.0.grid_primary", "infoblox.localdomain"),
					resource.TestCheckResourceAttr(resourceName, "member_soa_mnames.0.mname", "infoblox.localdomain"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthMemberSoaMnames(zoneFqdn, "default", gridPrimary, updatedMemberSoaMnames),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "member_soa_mnames.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "member_soa_mnames.0.grid_primary", "infoblox.localdomain"),
					resource.TestCheckResourceAttr(resourceName, "member_soa_mnames.0.mname", "example.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_MsAdIntegrated(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_ms_ad_integrated"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthMsAdIntegrated(zoneFqdn, "default", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_ad_integrated", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthMsAdIntegrated(zoneFqdn, "default", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_ad_integrated", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_MsAllowTransfer(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_ms_allow_transfer"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"
	msAllowTransfer := []map[string]any{
		{
			"address":    "192.168.1.10",
			"permission": "ALLOW",
		},
	}
	updatedMsAllowTransfer := []map[string]any{
		{
			"address":    "192.168.1.20",
			"permission": "DENY",
		},
	}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthMsAllowTransfer(zoneFqdn, "default", msAllowTransfer),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_allow_transfer.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ms_allow_transfer.0.address", "192.168.1.10"),
					resource.TestCheckResourceAttr(resourceName, "ms_allow_transfer.0.permission", "ALLOW"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthMsAllowTransfer(zoneFqdn, "default", updatedMsAllowTransfer),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_allow_transfer.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ms_allow_transfer.0.address", "192.168.1.20"),
					resource.TestCheckResourceAttr(resourceName, "ms_allow_transfer.0.permission", "DENY"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_MsAllowTransferMode(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_ms_allow_transfer_mode"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthMsAllowTransferMode(zoneFqdn, "default", "ANY"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_allow_transfer_mode", "ANY"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthMsAllowTransferMode(zoneFqdn, "default", "NONE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_allow_transfer_mode", "NONE"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_MsDcNsRecordCreation(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_ms_dc_ns_record_creation"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"
	msDcRecordCreation := []map[string]any{
		{
			"address": "10.0.0.0",
		},
	}
	updatedMsDcRecordCreation := []map[string]any{
		{
			"address": "198.51.100.0",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthMsDcNsRecordCreation(zoneFqdn, "default", msDcRecordCreation, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_dc_ns_record_creation.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ms_dc_ns_record_creation.0.address", "10.0.0.0"),
					resource.TestCheckResourceAttr(resourceName, "ms_ad_integrated", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthMsDcNsRecordCreation(zoneFqdn, "default", updatedMsDcRecordCreation, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_dc_ns_record_creation.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ms_dc_ns_record_creation.0.address", "198.51.100.0"),
					resource.TestCheckResourceAttr(resourceName, "ms_ad_integrated", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_MsDdnsMode(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_ms_ddns_mode"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthMsDdnsMode(zoneFqdn, "default", "ANY"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_ddns_mode", "ANY"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthMsDdnsMode(zoneFqdn, "default", "NONE"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_ddns_mode", "NONE"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_MsPrimaries(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_ms_primaries"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"
	msPrimaries := []map[string]any{
		{
			"address": "10.120.23.22",
			"ns_name": "example-server",
			"ns_ip":   "10.120.23.22",
		},
	}
	updatedMsPrimaries := []map[string]any{
		{
			"address": "10.120.23.23",
			"ns_name": "example-server",
			"ns_ip":   "10.120.23.23",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthMsPrimaries(zoneFqdn, "default", msPrimaries),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_primaries.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ms_primaries.0.address", "10.120.23.22"),
					resource.TestCheckResourceAttr(resourceName, "ms_primaries.0.ns_name", "example-server"),
					resource.TestCheckResourceAttr(resourceName, "ms_primaries.0.ns_ip", "10.120.23.22"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthMsPrimaries(zoneFqdn, "default", updatedMsPrimaries),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_primaries.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ms_primaries.0.address", "10.120.23.23"),
					resource.TestCheckResourceAttr(resourceName, "ms_primaries.0.ns_name", "example-server"),
					resource.TestCheckResourceAttr(resourceName, "ms_primaries.0.ns_ip", "10.120.23.23"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_MsSecondaries(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_ms_secondaries"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"
	msSecondaries := []map[string]any{
		{
			"address": "10.120.23.22",
			"ns_name": "example-server",
			"ns_ip":   "10.120.23.22",
		},
	}
	updatedMsSecondaries := []map[string]any{
		{
			"address": "10.120.23.23",
			"ns_name": "example-server",
			"ns_ip":   "10.120.23.23",
		},
	}
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
				Config: testAccZoneAuthMsSecondaries(zoneFqdn, "default", msSecondaries, gridPrimary),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_secondaries.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ms_secondaries.0.address", "10.120.23.22"),
					resource.TestCheckResourceAttr(resourceName, "ms_secondaries.0.ns_name", "example-server"),
					resource.TestCheckResourceAttr(resourceName, "ms_secondaries.0.ns_ip", "10.120.23.22"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthMsSecondaries(zoneFqdn, "default", updatedMsSecondaries, gridPrimary),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_secondaries.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "ms_secondaries.0.address", "10.120.23.23"),
					resource.TestCheckResourceAttr(resourceName, "ms_secondaries.0.ns_name", "example-server"),
					resource.TestCheckResourceAttr(resourceName, "ms_secondaries.0.ns_ip", "10.120.23.23"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_MsSyncDisabled(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_ms_sync_disabled"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"
	msPrimaries := []map[string]any{
		{
			"address": "10.120.23.22",
			"ns_name": "example-server",
			"ns_ip":   "10.120.23.22",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthMsSyncDisabled(zoneFqdn, "default", msPrimaries, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_sync_disabled", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthMsSyncDisabled(zoneFqdn, "default", msPrimaries, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_sync_disabled", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_NotifyDelay(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_notify_delay"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthNotifyDelay(zoneFqdn, "default", 5, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "notify_delay", "5"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthNotifyDelay(zoneFqdn, "default", 20, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "notify_delay", "20"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_NsGroup(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_ns_group"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthNsGroup(zoneFqdn, "default", "example-ns-group"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ns_group", "example-ns-group"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthNsGroup(zoneFqdn, "default", "updated-example-ns-group"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ns_group", "updated-example-ns-group"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_Prefix(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_prefix"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthPrefix(zoneFqdn, "default", "128/26"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "prefix", "128/26"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthPrefix(zoneFqdn, "default", "121/26"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "prefix", "121/26"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_RecordNamePolicy(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_record_name_policy"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthRecordNamePolicy(zoneFqdn, "default", "example-policy", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "record_name_policy", "example-policy"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthRecordNamePolicy(zoneFqdn, "default", "example-policy-update", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "record_name_policy", "example-policy-update"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_ScavengingSettings(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_scavenging_settings"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"
	scavengingSettings := map[string]any{
		"enable_scavenging": true,
		"expression_list": []map[string]any{
			{
				"op":       "AND",
				"op1_type": "LIST",
			},
			{
				"op":       "EQ",
				"op1":      "rtype",
				"op1_type": "FIELD",
				"op2":      "A",
				"op2_type": "STRING",
			},
			{
				"op": "ENDLIST",
			},
		},
	}
	updatedScavengingSettings := map[string]any{
		"enable_scavenging": true,
		"expression_list": []map[string]any{
			{
				"op":       "AND",
				"op1_type": "LIST",
			},
			{
				"op":       "EQ",
				"op1":      "rtype",
				"op1_type": "FIELD",
				"op2":      "AAAA",
				"op2_type": "STRING",
			},
			{
				"op": "ENDLIST",
			},
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthScavengingSettings(zoneFqdn, "default", scavengingSettings, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.enable_scavenging", "true"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.expression_list.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.expression_list.0.op", "AND"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.expression_list.0.op1_type", "LIST"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.expression_list.1.op", "EQ"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.expression_list.1.op1", "rtype"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.expression_list.1.op1_type", "FIELD"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.expression_list.1.op2", "A"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.expression_list.1.op2_type", "STRING"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.expression_list.2.op", "ENDLIST"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthScavengingSettings(zoneFqdn, "default", updatedScavengingSettings, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.enable_scavenging", "true"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.expression_list.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.expression_list.0.op", "AND"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.expression_list.0.op1_type", "LIST"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.expression_list.1.op", "EQ"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.expression_list.1.op1", "rtype"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.expression_list.1.op1_type", "FIELD"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.expression_list.1.op2", "AAAA"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.expression_list.1.op2_type", "STRING"),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings.expression_list.2.op", "ENDLIST"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_SoaDefaultTtl(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_soa_default_ttl"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"
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
				Config: testAccZoneAuthSoaDefaultTtl(zoneFqdn, "default", gridPrimary, 8, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_default_ttl", "8"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthSoaDefaultTtl(zoneFqdn, "default", gridPrimary, 10, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_default_ttl", "10"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_SoaEmail(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_soa_email"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"
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
				Config: testAccZoneAuthSoaEmail(zoneFqdn, "default", gridPrimary, "user1@example.com", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_email", "user1@example.com"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthSoaEmail(zoneFqdn, "default", gridPrimary, "user2@example.com", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_email", "user2@example.com"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_SoaExpire(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_soa_expire"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"
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
				Config: testAccZoneAuthSoaExpire(zoneFqdn, "default", gridPrimary, 24192, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_expire", "24192"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthSoaExpire(zoneFqdn, "default", gridPrimary, 24100, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_expire", "24100"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_SoaNegativeTtl(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_soa_negative_ttl"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"
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
				Config: testAccZoneAuthSoaNegativeTtl(zoneFqdn, "default", gridPrimary, 800, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_negative_ttl", "800"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthSoaNegativeTtl(zoneFqdn, "default", gridPrimary, 900, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_negative_ttl", "900"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_SoaRefresh(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_soa_refresh"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"
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
				Config: testAccZoneAuthSoaRefresh(zoneFqdn, "default", gridPrimary, 800, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_refresh", "800"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthSoaRefresh(zoneFqdn, "default", gridPrimary, 900, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_refresh", "900"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_SoaRetry(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_soa_retry"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"
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
				Config: testAccZoneAuthSoaRetry(zoneFqdn, "default", gridPrimary, 1600, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_retry", "1600"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthSoaRetry(zoneFqdn, "default", gridPrimary, 1700, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_retry", "1700"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_SoaSerialNumber(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_soa_serial_number"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"
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
				Config: testAccZoneAuthSoaSerialNumber(zoneFqdn, "default", gridPrimary, 10, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_serial_number", "10"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthSoaSerialNumber(zoneFqdn, "default", gridPrimary, 20, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_serial_number", "20"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_Srgs(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_srgs"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"
	sharedRecordGroup := []string{"example-shared-record-group"}
	updatedSharedRecordGroup := []string{"updated-shared-record-group"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthSrgs(zoneFqdn, "default", sharedRecordGroup),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "srgs.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "srgs.0", "example-shared-record-group"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthSrgs(zoneFqdn, "default", updatedSharedRecordGroup),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "srgs.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "srgs.0", "updated-shared-record-group"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_UpdateForwarding(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_update_forwarding"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"
	updateForwarding := []map[string]any{
		{
			"struct":     "addressac",
			"address":    "10.0.0.0",
			"permission": "ALLOW",
		},
	}
	updatedUpdateForwarding := []map[string]any{
		{
			"struct":        "tsigac",
			"tsig_key":      "X4oRe92t54I+T98NdQpV2w==",
			"tsig_key_name": "example-tsig-key",
			"tsig_key_alg":  "HMAC-SHA256",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthUpdateForwarding(zoneFqdn, "default", updateForwarding, true, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_forwarding.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "update_forwarding.0.address", "10.0.0.0"),
					resource.TestCheckResourceAttr(resourceName, "update_forwarding.0.permission", "ALLOW"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthUpdateForwarding(zoneFqdn, "default", updatedUpdateForwarding, true, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_forwarding.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "update_forwarding.0.tsig_key", "X4oRe92t54I+T98NdQpV2w=="),
					// resource.TestCheckResourceAttr(resourceName, "update_forwarding.0.permission", "ALLOW"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_UseAllowActiveDir(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_use_allow_active_dir"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"
	allowActiveDir := []map[string]any{
		{
			"address":    "10.0.0.1",
			"permission": "ALLOW",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthUseAllowActiveDir(zoneFqdn, "default", allowActiveDir, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_allow_active_dir", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthUseAllowActiveDir(zoneFqdn, "default", allowActiveDir, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_allow_active_dir", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_UseAllowQuery(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_use_allow_query"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"
	allowQuery := []map[string]any{
		{
			"struct":     "addressac",
			"address":    "10.0.0.0",
			"permission": "ALLOW",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthUseAllowQuery(zoneFqdn, "default", allowQuery, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_allow_query", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthUseAllowQuery(zoneFqdn, "default", allowQuery, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_allow_query", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_UseAllowTransfer(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_use_allow_transfer"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"
	allowTransfer := []map[string]any{
		{
			"struct":     "addressac",
			"address":    "10.0.0.0",
			"permission": "ALLOW",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthUseAllowTransfer(zoneFqdn, "default", allowTransfer, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_allow_transfer", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthUseAllowTransfer(zoneFqdn, "default", allowTransfer, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_allow_transfer", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_UseAllowUpdate(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_use_allow_update"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"
	allowUpdate := []map[string]any{
		{
			"struct":     "addressac",
			"address":    "10.0.0.0",
			"permission": "ALLOW",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthUseAllowUpdate(zoneFqdn, "default", allowUpdate, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_allow_update", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthUseAllowUpdate(zoneFqdn, "default", allowUpdate, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_allow_update", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_UseAllowUpdateForwarding(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_use_allow_update_forwarding"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"
	updateForwarding := []map[string]any{
		{
			"struct":     "addressac",
			"address":    "10.0.0.0",
			"permission": "ALLOW",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthUseAllowUpdateForwarding(zoneFqdn, "default", updateForwarding, true, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_allow_update_forwarding", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthUseAllowUpdateForwarding(zoneFqdn, "default", updateForwarding, true, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_allow_update_forwarding", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_UseCheckNamesPolicy(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_use_check_names_policy"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthUseCheckNamesPolicy(zoneFqdn, "default", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_check_names_policy", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthUseCheckNamesPolicy(zoneFqdn, "default", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_check_names_policy", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_UseCopyXferToNotify(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_use_copy_xfer_to_notify"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthUseCopyXferToNotify(zoneFqdn, "default", true, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_copy_xfer_to_notify", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthUseCopyXferToNotify(zoneFqdn, "default", true, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_copy_xfer_to_notify", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_UseDdnsForceCreationTimestampUpdate(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_use_ddns_force_creation_timestamp_update"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthUseDdnsForceCreationTimestampUpdate(zoneFqdn, "default", true, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_force_creation_timestamp_update", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthUseDdnsForceCreationTimestampUpdate(zoneFqdn, "default", true, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_force_creation_timestamp_update", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_UseDdnsPatternsRestriction(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_use_ddns_patterns_restriction"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthUseDdnsPatternsRestriction(zoneFqdn, "default", true, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_patterns_restriction", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthUseDdnsPatternsRestriction(zoneFqdn, "default", true, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_patterns_restriction", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_UseDdnsPrincipalSecurity(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_use_ddns_principal_security"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthUseDdnsPrincipalSecurity(zoneFqdn, "default", true, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_principal_security", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthUseDdnsPrincipalSecurity(zoneFqdn, "default", true, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_principal_security", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_UseDdnsRestrictProtected(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_use_ddns_restrict_protected"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthUseDdnsRestrictProtected(zoneFqdn, "default", true, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_restrict_protected", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthUseDdnsRestrictProtected(zoneFqdn, "default", true, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_restrict_protected", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_UseDdnsRestrictStatic(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_use_ddns_restrict_static"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthUseDdnsRestrictStatic(zoneFqdn, "default", true, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_restrict_static", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthUseDdnsRestrictStatic(zoneFqdn, "default", true, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_restrict_static", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_UseDnssecKeyParams(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_use_dnssec_key_params"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"
	kskAlgorithms := []map[string]any{
		{
			"algorithm": "RSASHA256",
			"size":      2048,
		},
	}
	zskAlgorithms := []map[string]any{
		{
			"algorithm": "RSASHA256",
			"size":      1024,
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthUseDnssecKeyParams(zoneFqdn, "default", kskAlgorithms, zskAlgorithms, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_dnssec_key_params", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthUseDnssecKeyParams(zoneFqdn, "default", kskAlgorithms, zskAlgorithms, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_dnssec_key_params", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_UseExternalPrimary(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_use_external_primary"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"
	externalPrimaries := []map[string]any{
		{
			"address": "10.0.0.0",
			"name":    "example-server",
		},
	}
	msSecondaries := []map[string]any{
		{
			"address": "10.120.23.22",
			"ns_name": "example-server",
			"ns_ip":   "10.120.23.22",
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthUseExternalPrimary(zoneFqdn, "default", externalPrimaries, msSecondaries, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_external_primary", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthUseExternalPrimaryUpdate(zoneFqdn, "default", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_external_primary", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_UseGridZoneTimer(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_use_grid_zone_timer"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"
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
				Config: testAccZoneAuthUseGridZoneTimer(zoneFqdn, "default", gridPrimary, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_grid_zone_timer", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthUseGridZoneTimerUpdate(zoneFqdn, "default", gridPrimary, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_grid_zone_timer", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_UseNotifyDelay(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_use_notify_delay"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthUseNotifyDelay(zoneFqdn, "default", 5, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_notify_delay", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthUseNotifyDelay(zoneFqdn, "default", 5, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_notify_delay", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_UseRecordNamePolicy(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_use_record_name_policy"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthUseRecordNamePolicy(zoneFqdn, "default", "example-policy", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_record_name_policy", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthUseRecordNamePolicy(zoneFqdn, "default", "example-policy", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_record_name_policy", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_UseScavengingSettings(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_use_scavenging_settings"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"
	scavengingSettings := map[string]any{
		"enable_scavenging": true,
		"expression_list": []map[string]any{
			{
				"op":       "AND",
				"op1_type": "LIST",
			},
			{
				"op":       "EQ",
				"op1":      "rtype",
				"op1_type": "FIELD",
				"op2":      "A",
				"op2_type": "STRING",
			},
			{
				"op": "ENDLIST",
			},
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthUseScavengingSettings(zoneFqdn, "default", scavengingSettings, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_scavenging_settings", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthUseScavengingSettings(zoneFqdn, "default", scavengingSettings, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_scavenging_settings", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_UseSoaEmail(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_use_soa_email"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"
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
				Config: testAccZoneAuthUseSoaEmail(zoneFqdn, "default", gridPrimary, "user@example.com", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_soa_email", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthUseSoaEmail(zoneFqdn, "default", gridPrimary, "user@example.com", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_soa_email", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_View(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_view"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthView(zoneFqdn, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthView(zoneFqdn, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", "default"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_ZoneFormatIPV4(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_zone_format"
	var v dns.ZoneAuth
	zoneFqdn := "10.0.0.0/24"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthZoneFormat(zoneFqdn, "default", "IPV4"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "zone_format", "IPV4"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_ZoneFormatIPV6(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_zone_format"
	var v dns.ZoneAuth
	zoneFqdn := "2002:1100::/64"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthZoneFormat(zoneFqdn, "default", "IPV6"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "zone_format", "IPV6"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckZoneAuthExists(ctx context.Context, resourceName string, v *dns.ZoneAuth) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DNSAPI.
			ZoneAuthAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForZoneAuth).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetZoneAuthResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetZoneAuthResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckZoneAuthDestroy(ctx context.Context, v *dns.ZoneAuth) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DNSAPI.
			ZoneAuthAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForZoneAuth).
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

func testAccCheckZoneAuthDisappears(ctx context.Context, v *dns.ZoneAuth) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DNSAPI.
			ZoneAuthAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccZoneAuthBasicConfig(zoneFqdn, view string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test" {
    fqdn = %q
    view = %q
}
`, zoneFqdn, view)
}

func testAccZoneAuthAllowActiveDir(zoneFqdn, view string, allowActiveDir []map[string]any, useAllowActiveDir bool) string {
	allowActiveDirHCL := utils.ConvertSliceOfMapsToHCL(allowActiveDir)
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_allow_active_dir" {
    fqdn = %q
    view = %q
    use_allow_active_dir = %t
    allow_active_dir = %s
}
`, zoneFqdn, view, useAllowActiveDir, allowActiveDirHCL)
}

func testAccZoneAuthAllowFixedRrsetOrder(zoneFqdn, view string, allowFixedRrsetOrder bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_allow_fixed_rrset_order" {
    fqdn = %q
    view = %q
    allow_fixed_rrset_order = %t
}
`, zoneFqdn, view, allowFixedRrsetOrder)
}

func testAccZoneAuthAllowGssTsigForUnderscoreZone(zoneFqdn, view string, allowGssTsigForUnderscoreZone bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_allow_gss_tsig_for_underscore_zone" {
    fqdn = %q
    view = %q
    allow_gss_tsig_for_underscore_zone = %t
}
`, zoneFqdn, view, allowGssTsigForUnderscoreZone)
}

func testAccZoneAuthAllowGssTsigZoneUpdates(zoneFqdn, view string, allowGssTsigZoneUpdates bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_allow_gss_tsig_zone_updates" {
    fqdn = %q
    view = %q
    allow_gss_tsig_zone_updates = %t
	allow_fixed_rrset_order = false
}
`, zoneFqdn, view, allowGssTsigZoneUpdates)
}

func testAccZoneAuthAllowQuery(zoneFqdn, view string, allowQuery []map[string]any, useAllowQuery bool) string {
	allowQueryHCL := utils.ConvertSliceOfMapsToHCL(allowQuery)
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_allow_query" {
    fqdn = %q
    view = %q
    allow_query = %s
	use_allow_query = %t
}
`, zoneFqdn, view, allowQueryHCL, useAllowQuery)
}

func testAccZoneAuthAllowTransfer(zoneFqdn, view string, allowTransfer []map[string]any, useAllowTransfer bool) string {
	allowTransferHCL := utils.ConvertSliceOfMapsToHCL(allowTransfer)
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_allow_transfer" {
    fqdn = %q
    view = %q
    allow_transfer = %s
	use_allow_transfer = %t
}
`, zoneFqdn, view, allowTransferHCL, useAllowTransfer)
}

func testAccZoneAuthAllowUpdate(zoneFqdn, view string, allowUpdate []map[string]any, useAllowUpdate bool) string {
	allowUpdateHCL := utils.ConvertSliceOfMapsToHCL(allowUpdate)
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_allow_update" {
    fqdn = %q
    view = %q
    allow_update = %s
    use_allow_update = %t
}
`, zoneFqdn, view, allowUpdateHCL, useAllowUpdate)
}

func testAccZoneAuthAllowUpdateForwarding(zoneFqdn, view string, allowUpdateForwarding, useAllowUpdateForwarding bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_allow_update_forwarding" {
    fqdn = %q
    view = %q
    allow_update_forwarding = %t
    use_allow_update_forwarding = %t
}
`, zoneFqdn, view, allowUpdateForwarding, useAllowUpdateForwarding)
}
func testAccZoneAuthComment(zoneFqdn, view, comment string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_comment" {
    fqdn = %q
    view = %q
    comment = %q
}
`, zoneFqdn, view, comment)
}

func testAccZoneAuthCopyXferToNotify(zoneFqdn, view string, copyXferToNotify, useCopyXferToNotify bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_copy_xfer_to_notify" {
    fqdn = %q
    view = %q
    copy_xfer_to_notify = %t
    use_copy_xfer_to_notify = %t
}
`, zoneFqdn, view, copyXferToNotify, useCopyXferToNotify)
}

func testAccZoneAuthCreateUnderscoreZones(zoneFqdn, view string, createUnderscoreZones bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_create_underscore_zones" {
    fqdn = %q
    view = %q
    create_underscore_zones = %t
}
`, zoneFqdn, view, createUnderscoreZones)
}

func testAccZoneAuthDdnsForceCreationTimestampUpdate(zoneFqdn, view string, ddnsForceCreationTimestampUpdate, useDdnsForceCreationTimestampUpdate bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_ddns_force_creation_timestamp_update" {
	fqdn = %q
	view = %q
    ddns_force_creation_timestamp_update = %t
	use_ddns_force_creation_timestamp_update = %t
}
`, zoneFqdn, view, ddnsForceCreationTimestampUpdate, useDdnsForceCreationTimestampUpdate)
}

func testAccZoneAuthDdnsPrincipalGroup(zoneFqdn, view, ddnsPrincipalGroup string, useDdnsPrincipalSecurity bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_ddns_principal_group" {
    fqdn = %q
    view = %q
    ddns_principal_group = %q
	use_ddns_principal_security = %t
}
`, zoneFqdn, view, ddnsPrincipalGroup, useDdnsPrincipalSecurity)
}

func testAccZoneAuthDdnsPrincipalTracking(zoneFqdn, view string, ddnsPrincipalTracking, useDdnsPrincipalSecurity bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_ddns_principal_tracking" {
    fqdn = %q
    view = %q
    ddns_principal_tracking = %t
	use_ddns_principal_security = %t
}
`, zoneFqdn, view, ddnsPrincipalTracking, useDdnsPrincipalSecurity)
}

func testAccZoneAuthDdnsRestrictPatterns(zoneFqdn, view string, ddnsRestrictPatterns, useDdnsRestrictPatternRestriction bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_ddns_restrict_patterns" {
    fqdn = %q
    view = %q
    ddns_restrict_patterns = %t
    use_ddns_patterns_restriction = %t
}
`, zoneFqdn, view, ddnsRestrictPatterns, useDdnsRestrictPatternRestriction)
}

func testAccZoneAuthDdnsRestrictPatternsList(zoneFqdn, view string, ddnsRestrictPatternsList []string, useDdnsRestrictPatternRestriction bool) string {
	patterns := utils.ConvertStringSliceToHCL(ddnsRestrictPatternsList)

	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_ddns_restrict_patterns_list" {
    fqdn = %q
    view = %q
    ddns_restrict_patterns_list = %s
    use_ddns_patterns_restriction = %t
}
`, zoneFqdn, view, patterns, useDdnsRestrictPatternRestriction)
}

func testAccZoneAuthDdnsRestrictProtected(zoneFqdn, view string, ddnsRestrictProtected, useDdnsRestrictProtected bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_ddns_restrict_protected" {
    fqdn = %q
    view = %q
    ddns_restrict_protected = %t
    use_ddns_restrict_protected = %t
}
`, zoneFqdn, view, ddnsRestrictProtected, useDdnsRestrictProtected)
}

func testAccZoneAuthDdnsRestrictSecure(zoneFqdn, view string, ddnsRestrictSecure, useDdnsPrincipalSecurity bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_ddns_restrict_secure" {
    fqdn = %q
    view = %q
    ddns_restrict_secure = %t
    use_ddns_principal_security = %t
}
`, zoneFqdn, view, ddnsRestrictSecure, useDdnsPrincipalSecurity)
}

func testAccZoneAuthDdnsRestrictStatic(zoneFqdn, view string, ddnsRestrictStatic, useDdnsRestrictStatic bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_ddns_restrict_static" {
    fqdn = %q
    view = %q
    ddns_restrict_static = %t
    use_ddns_restrict_static = %t
}
`, zoneFqdn, view, ddnsRestrictStatic, useDdnsRestrictStatic)
}

func testAccZoneAuthDisable(zoneFqdn, view string, disable bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_disable" {
    fqdn = %q
    view = %q
    disable = %t
}
`, zoneFqdn, view, disable)
}

func testAccZoneAuthDisableForwarding(zoneFqdn, view string, disableForwarding bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_disable_forwarding" {
    fqdn = %q
    view = %q
    disable_forwarding = %t
}
`, zoneFqdn, view, disableForwarding)
}

func testAccZoneAuthDnsIntegrityEnable(zoneFqdn, view string, dnsIntegrityEnable bool, gridPrimary []map[string]any, dnsIntegrityMember string) string {
	gridPrimaryHCL := utils.ConvertSliceOfMapsToHCL(gridPrimary)
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_dns_integrity_enable" {
    fqdn = %q
    view = %q
    dns_integrity_enable = %t
    dns_integrity_member = %q
	grid_primary = %s
}`, zoneFqdn, view, dnsIntegrityEnable, dnsIntegrityMember, gridPrimaryHCL)
}

func testAccZoneAuthDnsIntegrityFrequency(zoneFqdn, view string, dnsIntegrityFrequency int) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_dns_integrity_frequency" {
    fqdn = %q
    view = %q
    dns_integrity_frequency = %d
}
`, zoneFqdn, view, dnsIntegrityFrequency)
}

func testAccZoneAuthDnsIntegrityMember(zoneFqdn, view string, dnsIntegrityMember string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_dns_integrity_member" {
    fqdn = %q
    view = %q
    dns_integrity_member = %q
}
`, zoneFqdn, view, dnsIntegrityMember)
}

func testAccZoneAuthDnsIntegrityVerboseLogging(zoneFqdn, view string, dnsIntegrityVerboseLogging bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_dns_integrity_verbose_logging" {
    fqdn = %q
    view = %q
    dns_integrity_verbose_logging = %t
}
`, zoneFqdn, view, dnsIntegrityVerboseLogging)
}

func testAccZoneAuthDnssecKeyParams(zoneFqdn, view string, kskAlgorithms, zskAlgorithms []map[string]any) string {
	// Use the helper functions to build configurations
	kskHCL := utils.ConvertSliceOfMapsToHCL(kskAlgorithms)
	zskCHCL := utils.ConvertSliceOfMapsToHCL(zskAlgorithms)

	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_dnssec_key_params" {
  fqdn = %q
  view = %q
  use_dnssec_key_params = true
  dnssec_key_params = {
  ksk_algorithms = %s
  zsk_algorithms = %s
  }
}
`, zoneFqdn, view, kskHCL, zskCHCL)
}

func testAccZoneAuthEffectiveCheckNamesPolicy(zoneFqdn, view, effectiveCheckNamesPolicy string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_effective_check_names_policy" {
	fqdn = %q
    view = %q
    effective_check_names_policy = %q
}
`, zoneFqdn, view, effectiveCheckNamesPolicy)
}

func testAccZoneAuthExtAttrs(zoneFqdn, view string, extAttrs map[string]string) string {
	extattrsStr := "{\n"
	for k, v := range extAttrs {
		extattrsStr += fmt.Sprintf(`
  %s = %q
`, k, v)
	}
	extattrsStr += "\t}"
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_extattrs" {
    fqdn = %q
    view = %q
    extattrs = %s
}
`, zoneFqdn, view, extattrsStr)
}

func testAccZoneAuthExternalPrimaries(zoneFqdn, view string, externalPrimaries, msSecondaries []map[string]any, useExternalPrimary bool) string {
	externalPrimariesHCL := utils.ConvertSliceOfMapsToHCL(externalPrimaries)
	msSecondariesHCL := utils.ConvertSliceOfMapsToHCL(msSecondaries)
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_external_primaries" {
	fqdn = %q
    view = %q
    external_primaries = %s
	ms_secondaries = %s
	use_external_primary = %t
}`, zoneFqdn, view, externalPrimariesHCL, msSecondariesHCL, useExternalPrimary)
}

func testAccZoneAuthExternalSecondaries(zoneFqdn, view string, externalSecondaries, gridPrimary []map[string]any) string {
	externalSecondariesHCL := utils.ConvertSliceOfMapsToHCL(externalSecondaries)
	gridPrimaryHCL := utils.ConvertSliceOfMapsToHCL(gridPrimary)
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_external_secondaries" {
    fqdn = %q
    view = %q
    grid_primary = %s
    external_secondaries = %s
}
`, zoneFqdn, view, gridPrimaryHCL, externalSecondariesHCL)
}

func testAccZoneAuthFqdn(fqdn, view string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_fqdn" {
    fqdn = %q
    view = %q
}
`, fqdn, view)
}

func testAccZoneAuthGridPrimary(zoneFqdn, view string, gridPrimary []map[string]any) string {
	gridPrimaryHCL := utils.ConvertSliceOfMapsToHCL(gridPrimary)
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_grid_primary" {
    fqdn = %q
    view = %q
    grid_primary = %s
}
`, zoneFqdn, view, gridPrimaryHCL)
}

func testAccZoneAuthGridSecondaries(zoneFqdn, view string, gridPrimary, gridSecondary []map[string]any) string {
	gridPrimaryHCL := utils.ConvertSliceOfMapsToHCL(gridPrimary)
	gridSecondaryHCL := utils.ConvertSliceOfMapsToHCL(gridSecondary)
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_grid_secondaries" {
	fqdn = %q
	view = %q
	grid_primary = %s
    grid_secondaries = %s
}
`, zoneFqdn, view, gridPrimaryHCL, gridSecondaryHCL)
}

func testAccZoneAuthLastQueriedAcl(zoneFqdn, view string, lastQueriedACL []map[string]any, useScavengingSettings bool) string {
	lastQueriedACLAsHCL := utils.ConvertSliceOfMapsToHCL(lastQueriedACL)
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_last_queried_acl" {
    fqdn = %q
    view = %q
    last_queried_acl = %s
	use_scavenging_settings = %t
}
`, zoneFqdn, view, lastQueriedACLAsHCL, useScavengingSettings)
}

func testAccZoneAuthLocked(zoneFqdn, view string, locked bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_locked" {
    fqdn = %q
    view = %q
    locked = %t
}
`, zoneFqdn, view, locked)
}

func testAccZoneAuthMemberSoaMnames(zoneFqdn, view string, gridPrimary, memberSoaMnames []map[string]any) string {
	gridPrimaryHCL := utils.ConvertSliceOfMapsToHCL(gridPrimary)
	memberSoaMnamesHCL := utils.ConvertSliceOfMapsToHCL(memberSoaMnames)
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_member_soa_mnames" {
    fqdn = %q
    view = %q
    grid_primary = %s
    member_soa_mnames = %s
}
`, zoneFqdn, view, gridPrimaryHCL, memberSoaMnamesHCL)
}

func testAccZoneAuthMsAdIntegrated(zoneFqdn, view string, msAdIntegrated bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_ms_ad_integrated" {
    fqdn = %q
    view = %q
    ms_ad_integrated = %t
}
`, zoneFqdn, view, msAdIntegrated)
}

func testAccZoneAuthMsAllowTransfer(zoneFqdn, view string, msAllowTransfer []map[string]any) string {
	msAllowTransferConfig := utils.ConvertSliceOfMapsToHCL(msAllowTransfer)
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_ms_allow_transfer" {
    fqdn = %q
    view = %q
    ms_allow_transfer = %s
}
`, zoneFqdn, view, msAllowTransferConfig)
}

func testAccZoneAuthMsAllowTransferMode(zoneFqdn, view, msAllowTransferMode string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_ms_allow_transfer_mode" {
    fqdn = %q
    view = %q
    ms_allow_transfer_mode = %q
}
`, zoneFqdn, view, msAllowTransferMode)
}

func testAccZoneAuthMsDcNsRecordCreation(zoneFqdn, view string, msDcRecordCreation []map[string]any, msAdIntegrated bool) string {
	msDcRecordCreationHCL := utils.ConvertSliceOfMapsToHCL(msDcRecordCreation)
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_ms_dc_ns_record_creation" {
    fqdn = %q
    view = %q
    ms_dc_ns_record_creation = %s
	ms_ad_integrated = %t
}
`, zoneFqdn, view, msDcRecordCreationHCL, msAdIntegrated)
}

func testAccZoneAuthMsDdnsMode(zoneFqdn, view, msDdnsMode string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_ms_ddns_mode" {
    fqdn = %q
    view = %q
    ms_ddns_mode = %q
}
`, zoneFqdn, view, msDdnsMode)
}

func testAccZoneAuthMsPrimaries(zoneFqdn, view string, msPrimaries []map[string]any) string {
	msPrimariesHCL := utils.ConvertSliceOfMapsToHCL(msPrimaries)
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_ms_primaries" {
    fqdn = %q
    view = %q
    ms_primaries = %s
}
`, zoneFqdn, view, msPrimariesHCL)
}

func testAccZoneAuthMsSecondaries(zoneFqdn, view string, msSecondaries []map[string]any, gridPrimary []map[string]any) string {
	msSecondariesHCL := utils.ConvertSliceOfMapsToHCL(msSecondaries)
	gridPrimaryHCL := utils.ConvertSliceOfMapsToHCL(gridPrimary)
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_ms_secondaries" {
    fqdn = %q
    view = %q
    ms_secondaries = %s
    grid_primary = %s
}
`, zoneFqdn, view, msSecondariesHCL, gridPrimaryHCL)
}

func testAccZoneAuthMsSyncDisabled(zoneFqdn, view string, msPrimaries []map[string]any, msSyncDisabled bool) string {
	msPrimariesHCL := utils.ConvertSliceOfMapsToHCL(msPrimaries)
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_ms_sync_disabled" {
    fqdn = %q
    view = %q
    ms_sync_disabled = %t
    ms_primaries = %s
}
`, zoneFqdn, view, msSyncDisabled, msPrimariesHCL)
}

func testAccZoneAuthNotifyDelay(zoneFqdn, view string, notifyDelay int, useNotifyDelay bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_notify_delay" {
    fqdn = %q
    view = %q
    notify_delay = %d
	use_notify_delay = %t

}
`, zoneFqdn, view, notifyDelay, useNotifyDelay)
}

func testAccZoneAuthNsGroup(zoneFqdn, view, nsGroup string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_ns_group" {
    fqdn = %q
    view = %q
    ns_group = %q
}
`, zoneFqdn, view, nsGroup)
}

func testAccZoneAuthPrefix(zoneFqdn, view, prefix string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_prefix" {
    fqdn = %q
    view = %q
    prefix = %q
}
`, zoneFqdn, view, prefix)
}

func testAccZoneAuthRecordNamePolicy(zoneFqdn, view, recordNamePolicy string, useRecordNamePolicy bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_record_name_policy" {
    fqdn = %q
    view = %q
    record_name_policy = %q
	use_record_name_policy = %t
}
`, zoneFqdn, view, recordNamePolicy, useRecordNamePolicy)
}

func testAccZoneAuthScavengingSettings(zoneFqdn, view string, scavengingSettings map[string]any, useScavengingSettings bool) string {
	scavengingSettingsHCL := utils.ConvertMapToHCL(scavengingSettings)
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_scavenging_settings" {
    fqdn = %q
    view = %q
    scavenging_settings = %s
	use_scavenging_settings = %t
}
`, zoneFqdn, view, scavengingSettingsHCL, useScavengingSettings)
}

func testAccZoneAuthSoaDefaultTtl(zoneFqdn, view string, gridPrimary []map[string]any, soaDefaultTtl int64, useGridZoneTimer bool) string {
	gridPrimaryHCL := utils.ConvertSliceOfMapsToHCL(gridPrimary)
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_soa_default_ttl" {
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

func testAccZoneAuthSoaEmail(zoneFqdn, view string, gridPrimary []map[string]any, soaEmail string, useSoaEmail bool) string {
	gridPrimaryHCL := utils.ConvertSliceOfMapsToHCL(gridPrimary)
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_soa_email" {
    fqdn = %q
    view = %q
    grid_primary = %s
    soa_email = %q
    use_soa_email = %t
}
`, zoneFqdn, view, gridPrimaryHCL, soaEmail, useSoaEmail)
}

func testAccZoneAuthSoaExpire(zoneFqdn, view string, gridPrimary []map[string]any, soaExpire int64, useGridZoneTimer bool) string {
	gridPrimaryHCL := utils.ConvertSliceOfMapsToHCL(gridPrimary)
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_soa_expire" {
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

func testAccZoneAuthSoaNegativeTtl(zoneFqdn, view string, gridPrimary []map[string]any, soaNegativeTtl int64, useGridZoneTimer bool) string {
	gridPrimaryHCL := utils.ConvertSliceOfMapsToHCL(gridPrimary)
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_soa_negative_ttl" {
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

func testAccZoneAuthSoaRefresh(zoneFqdn, view string, gridPrimary []map[string]any, soaRefresh int64, useGridZoneTimer bool) string {
	gridPrimaryHCL := utils.ConvertSliceOfMapsToHCL(gridPrimary)
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_soa_refresh" {
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

func testAccZoneAuthSoaRetry(zoneFqdn, view string, gridPrimary []map[string]any, soaRetry int64, useGridZoneTimer bool) string {
	gridPrimaryHCL := utils.ConvertSliceOfMapsToHCL(gridPrimary)
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_soa_retry" {
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

func testAccZoneAuthSoaSerialNumber(zoneFqdn, view string, gridPrimary []map[string]any, soaSerialNumber int64, SetSoaSerialNumber bool) string {
	gridPrimaryHCL := utils.ConvertSliceOfMapsToHCL(gridPrimary)
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_soa_serial_number" {
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

func testAccZoneAuthSrgs(zoneFqdn, view string, srgs []string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_srgs" {
    fqdn = %q
    view = %q
    srgs = %q
}
`, zoneFqdn, view, srgs)
}

func testAccZoneAuthUpdateForwarding(zoneFqdn, view string, updateForwarding []map[string]any, allowUpdateForwarding, useAllowUpdateForwarding bool) string {
	updateForwardingHCL := utils.ConvertSliceOfMapsToHCL(updateForwarding)
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_update_forwarding" {
    fqdn = %q
    view = %q
    update_forwarding = %s
    allow_update_forwarding = %t
    use_allow_update_forwarding = %t
}
`, zoneFqdn, view, updateForwardingHCL, allowUpdateForwarding, useAllowUpdateForwarding)
}

func testAccZoneAuthUseAllowActiveDir(zoneFqdn, view string, allowActiveDir []map[string]any, useAllowActiveDir bool) string {
	allowActiveDirHCL := utils.ConvertSliceOfMapsToHCL(allowActiveDir)
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_use_allow_active_dir" {
    fqdn = %q
    view = %q
    allow_active_dir = %s
    use_allow_active_dir = %t
}
`, zoneFqdn, view, allowActiveDirHCL, useAllowActiveDir)
}

func testAccZoneAuthUseAllowQuery(zoneFqdn, view string, allowQuery []map[string]any, useAllowQuery bool) string {
	allowQueryHCL := utils.ConvertSliceOfMapsToHCL(allowQuery)
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_use_allow_query" {
    fqdn = %q
    view = %q
    allow_query = %s
    use_allow_query = %t
}
`, zoneFqdn, view, allowQueryHCL, useAllowQuery)
}

func testAccZoneAuthUseAllowTransfer(zoneFqdn, view string, allowTransfer []map[string]any, useAllowTransfer bool) string {
	allowTransferHCL := utils.ConvertSliceOfMapsToHCL(allowTransfer)
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_use_allow_transfer" {
    fqdn = %q
    view = %q
    allow_transfer = %s
    use_allow_transfer = %t
}
`, zoneFqdn, view, allowTransferHCL, useAllowTransfer)
}

func testAccZoneAuthUseAllowUpdate(zoneFqdn, view string, allowUpdate []map[string]any, useAllowUpdate bool) string {
	allowUpdateHCL := utils.ConvertSliceOfMapsToHCL(allowUpdate)
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_use_allow_update" {
    fqdn = %q
    view = %q
    allow_update = %s
    use_allow_update = %t
}
`, zoneFqdn, view, allowUpdateHCL, useAllowUpdate)
}

func testAccZoneAuthUseAllowUpdateForwarding(zoneFqdn, view string, updateForwarding []map[string]any, allowUpdateForwarding, useAllowUpdateForwarding bool) string {
	updateForwardingHCL := utils.ConvertSliceOfMapsToHCL(updateForwarding)
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_use_allow_update_forwarding" {
    fqdn = %q
    view = %q
    update_forwarding = %s
    allow_update_forwarding = %t
    use_allow_update_forwarding = %t
}
`, zoneFqdn, view, updateForwardingHCL, allowUpdateForwarding, useAllowUpdateForwarding)
}

func testAccZoneAuthUseCheckNamesPolicy(zoneFqdn, view string, useCheckNamesPolicy bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_use_check_names_policy" {
    fqdn = %q
    view = %q
    use_check_names_policy = %t
}
`, zoneFqdn, view, useCheckNamesPolicy)
}

func testAccZoneAuthUseCopyXferToNotify(zoneFqdn, view string, copyXfer, useCopyXferToNotify bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_use_copy_xfer_to_notify" {
    fqdn = %q
    view = %q
    copy_xfer_to_notify = %t
    use_copy_xfer_to_notify = %t

}
`, zoneFqdn, view, copyXfer, useCopyXferToNotify)
}

func testAccZoneAuthUseDdnsForceCreationTimestampUpdate(zoneFqdn, view string, ddnsForceCreationTimestampUpdate, useDdnsForceCreationTimestampUpdate bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_use_ddns_force_creation_timestamp_update" {
    fqdn = %q
    view = %q
    ddns_force_creation_timestamp_update = %t
    use_ddns_force_creation_timestamp_update = %t
}
`, zoneFqdn, view, ddnsForceCreationTimestampUpdate, useDdnsForceCreationTimestampUpdate)
}

func testAccZoneAuthUseDdnsPatternsRestriction(zoneFqdn, view string, ddnsRestrictPatterns, useDdnsPatternsRestriction bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_use_ddns_patterns_restriction" {
    fqdn = %q
    view = %q
    ddns_restrict_patterns = %t
    use_ddns_patterns_restriction = %t
}
`, zoneFqdn, view, ddnsRestrictPatterns, useDdnsPatternsRestriction)
}

func testAccZoneAuthUseDdnsPrincipalSecurity(zoneFqdn, view string, ddnsRestrictSecure, useDdnsPrincipalSecurity bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_use_ddns_principal_security" {
    fqdn = %q
    view = %q
    ddns_restrict_secure = %t
    use_ddns_principal_security = %t
}
`, zoneFqdn, view, ddnsRestrictSecure, useDdnsPrincipalSecurity)
}

func testAccZoneAuthUseDdnsRestrictProtected(zoneFqdn, view string, ddnsRestrictProtected, useDdnsRestrictProtected bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_use_ddns_restrict_protected" {
    fqdn = %q
    view = %q
    ddns_restrict_protected = %t
    use_ddns_restrict_protected = %t
}
`, zoneFqdn, view, ddnsRestrictProtected, useDdnsRestrictProtected)
}

func testAccZoneAuthUseDdnsRestrictStatic(zoneFqdn, view string, ddnsRestrictStatic, useDdnsRestrictStatic bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_use_ddns_restrict_static" {
    fqdn = %q
    view = %q
    ddns_restrict_static = %t
    use_ddns_restrict_static = %t
}
`, zoneFqdn, view, ddnsRestrictStatic, useDdnsRestrictStatic)
}

func testAccZoneAuthUseDnssecKeyParams(zoneFqdn, view string, kskAlgorithms, zskAlgorithms []map[string]any, useDnssecKeyParams bool) string {
	kskAlgorithmsHCL := utils.ConvertSliceOfMapsToHCL(kskAlgorithms)
	zskAlgorithmsHCL := utils.ConvertSliceOfMapsToHCL(zskAlgorithms)
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_use_dnssec_key_params" {
    fqdn = %q
    view = %q
    dnssec_key_params = {
        ksk_algorithms = %s
        zsk_algorithms = %s
    }
	use_dnssec_key_params = %t
}`, zoneFqdn, view, kskAlgorithmsHCL, zskAlgorithmsHCL, useDnssecKeyParams)
}

func testAccZoneAuthUseExternalPrimary(zoneFqdn, view string, externalPrimaries, msSecondaries []map[string]any, useExternalPrimary bool) string {
	externalPrimariesHCL := utils.ConvertSliceOfMapsToHCL(externalPrimaries)
	msSecondariesHCL := utils.ConvertSliceOfMapsToHCL(msSecondaries)
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_use_external_primary" {
    fqdn = %q
    view = %q
    external_primaries = %s
    ms_secondaries = %s
    use_external_primary = %t
}
`, zoneFqdn, view, externalPrimariesHCL, msSecondariesHCL, useExternalPrimary)
}

func testAccZoneAuthUseExternalPrimaryUpdate(zoneFqdn, view string, useExternalPrimary bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_use_external_primary" {
    fqdn = %q
    view = %q
    use_external_primary = %t
}
`, zoneFqdn, view, useExternalPrimary)
}

func testAccZoneAuthUseGridZoneTimer(zoneFqdn, view string, gridPrimary []map[string]any, useGridZoneTimer bool) string {
	gridPrimaryHCL := utils.ConvertSliceOfMapsToHCL(gridPrimary)
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_use_grid_zone_timer" {
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

func testAccZoneAuthUseGridZoneTimerUpdate(zoneFqdn, view string, gridPrimary []map[string]any, useGridZoneTimer bool) string {
	gridPrimaryHCL := utils.ConvertSliceOfMapsToHCL(gridPrimary)
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_use_grid_zone_timer" {
    fqdn = %q
    view = %q
    grid_primary = %s
    use_grid_zone_timer = %t
}
`, zoneFqdn, view, gridPrimaryHCL, useGridZoneTimer)
}

func testAccZoneAuthUseNotifyDelay(zoneFqdn, view string, notifyDelay int64, useNotifyDelay bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_use_notify_delay" {
    fqdn = %q
    view = %q
    notify_delay = %d
    use_notify_delay = %t
}
`, zoneFqdn, view, notifyDelay, useNotifyDelay)
}

func testAccZoneAuthUseRecordNamePolicy(zoneFqdn, view, recordNamePolicy string, useRecordNamePolicy bool) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_use_record_name_policy" {
    fqdn = %q
    view = %q
    record_name_policy = %q
    use_record_name_policy = %t
}
`, zoneFqdn, view, recordNamePolicy, useRecordNamePolicy)
}

func testAccZoneAuthUseScavengingSettings(zoneFqdn, view string, scavengingSettings map[string]any, useScavengingSettings bool) string {
	scavengingSettingsHCL := utils.ConvertMapToHCL(scavengingSettings)
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_use_scavenging_settings" {
    fqdn = %q
    view = %q
    scavenging_settings = %s
    use_scavenging_settings = %t
}
`, zoneFqdn, view, scavengingSettingsHCL, useScavengingSettings)
}

func testAccZoneAuthUseSoaEmail(zoneFqdn, view string, gridPrimary []map[string]any, soaEmail string, useSoaEmail bool) string {
	gridPrimaryHCL := utils.ConvertSliceOfMapsToHCL(gridPrimary)
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_use_soa_email" {
    fqdn = %q
    view = %q
    grid_primary = %s
    soa_email = %q
    use_soa_email = %t
}
`, zoneFqdn, view, gridPrimaryHCL, soaEmail, useSoaEmail)
}

func testAccZoneAuthView(zoneFqdn, view string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_view" {
    fqdn = %q
    view = %q
}
`, zoneFqdn, view)
}

func testAccZoneAuthZoneFormat(zoneFqdn, view, zoneFormat string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_zone_format" {
    fqdn = %q
    view = %q
    zone_format = %q
}
`, zoneFqdn, view, zoneFormat)
}
