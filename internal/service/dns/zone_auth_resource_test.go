package dns_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dns"
	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/acctest"
	"github.com/Infoblox-CTO/infoblox-nios-terraform/internal/utils"
)

var readableAttributesForZoneAuth = "address,allow_active_dir,allow_fixed_rrset_order,allow_gss_tsig_for_underscore_zone,allow_gss_tsig_zone_updates,allow_query,allow_transfer,allow_update,allow_update_forwarding,aws_rte53_zone_info,cloud_info,comment,copy_xfer_to_notify,create_underscore_zones,ddns_force_creation_timestamp_update,ddns_principal_group,ddns_principal_tracking,ddns_restrict_patterns,ddns_restrict_patterns_list,ddns_restrict_protected,ddns_restrict_secure,ddns_restrict_static,disable,disable_forwarding,display_domain,dns_fqdn,dns_integrity_enable,dns_integrity_frequency,dns_integrity_member,dns_integrity_verbose_logging,dns_soa_email,dnssec_key_params,dnssec_keys,dnssec_ksk_rollover_date,dnssec_zsk_rollover_date,effective_check_names_policy,effective_record_name_policy,extattrs,external_primaries,external_secondaries,fqdn,grid_primary,grid_primary_shared_with_ms_parent_delegation,grid_secondaries,is_dnssec_enabled,is_dnssec_signed,is_multimaster,last_queried,last_queried_acl,locked,locked_by,mask_prefix,member_soa_mnames,member_soa_serials,ms_ad_integrated,ms_allow_transfer,ms_allow_transfer_mode,ms_dc_ns_record_creation,ms_ddns_mode,ms_managed,ms_primaries,ms_read_only,ms_secondaries,ms_sync_disabled,ms_sync_master_name,network_associations,network_view,notify_delay,ns_group,parent,prefix,primary_type,record_name_policy,records_monitored,rr_not_queried_enabled_time,scavenging_settings,soa_default_ttl,soa_email,soa_expire,soa_negative_ttl,soa_refresh,soa_retry,soa_serial_number,srgs,update_forwarding,use_allow_active_dir,use_allow_query,use_allow_transfer,use_allow_update,use_allow_update_forwarding,use_check_names_policy,use_copy_xfer_to_notify,use_ddns_force_creation_timestamp_update,use_ddns_patterns_restriction,use_ddns_principal_security,use_ddns_restrict_protected,use_ddns_restrict_static,use_dnssec_key_params,use_external_primary,use_grid_zone_timer,use_import_from,use_notify_delay,use_record_name_policy,use_scavenging_settings,use_soa_email,using_srg_associations,view,zone_format,zone_not_queried_enabled_time"

func TestAccZoneAuthResource_basic(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthBasicConfig(zoneFqdn, "default"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					// Test fields with default values
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_disappears(t *testing.T) {
	resourceName := "nios_dns_zone_auth.test"
	var v dns.ZoneAuth
	zoneFqdn := acctest.RandomNameWithPrefix("zone") + ".com"

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

func TestAccZoneAuthResource_Ref(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test__ref"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthRef("REF_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "_ref", "REF_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthRef("REF_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "_ref", "REF_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_AllowActiveDir(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_allow_active_dir"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthAllowActiveDir("ALLOW_ACTIVE_DIR_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_active_dir", "ALLOW_ACTIVE_DIR_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthAllowActiveDir("ALLOW_ACTIVE_DIR_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_active_dir", "ALLOW_ACTIVE_DIR_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_AllowFixedRrsetOrder(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_allow_fixed_rrset_order"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthAllowFixedRrsetOrder("ALLOW_FIXED_RRSET_ORDER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_fixed_rrset_order", "ALLOW_FIXED_RRSET_ORDER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthAllowFixedRrsetOrder("ALLOW_FIXED_RRSET_ORDER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_fixed_rrset_order", "ALLOW_FIXED_RRSET_ORDER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_AllowGssTsigForUnderscoreZone(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_allow_gss_tsig_for_underscore_zone"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthAllowGssTsigForUnderscoreZone("ALLOW_GSS_TSIG_FOR_UNDERSCORE_ZONE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_gss_tsig_for_underscore_zone", "ALLOW_GSS_TSIG_FOR_UNDERSCORE_ZONE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthAllowGssTsigForUnderscoreZone("ALLOW_GSS_TSIG_FOR_UNDERSCORE_ZONE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_gss_tsig_for_underscore_zone", "ALLOW_GSS_TSIG_FOR_UNDERSCORE_ZONE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_AllowGssTsigZoneUpdates(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_allow_gss_tsig_zone_updates"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthAllowGssTsigZoneUpdates("ALLOW_GSS_TSIG_ZONE_UPDATES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_gss_tsig_zone_updates", "ALLOW_GSS_TSIG_ZONE_UPDATES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthAllowGssTsigZoneUpdates("ALLOW_GSS_TSIG_ZONE_UPDATES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_gss_tsig_zone_updates", "ALLOW_GSS_TSIG_ZONE_UPDATES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_AllowQuery(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_allow_query"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthAllowQuery("ALLOW_QUERY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_query", "ALLOW_QUERY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthAllowQuery("ALLOW_QUERY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_query", "ALLOW_QUERY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_AllowTransfer(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_allow_transfer"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthAllowTransfer("ALLOW_TRANSFER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_transfer", "ALLOW_TRANSFER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthAllowTransfer("ALLOW_TRANSFER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_transfer", "ALLOW_TRANSFER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_AllowUpdate(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_allow_update"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthAllowUpdate("ALLOW_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_update", "ALLOW_UPDATE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthAllowUpdate("ALLOW_UPDATE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_update", "ALLOW_UPDATE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_AllowUpdateForwarding(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_allow_update_forwarding"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthAllowUpdateForwarding("ALLOW_UPDATE_FORWARDING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_update_forwarding", "ALLOW_UPDATE_FORWARDING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthAllowUpdateForwarding("ALLOW_UPDATE_FORWARDING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_update_forwarding", "ALLOW_UPDATE_FORWARDING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_AwsRte53ZoneInfo(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_aws_rte53_zone_info"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthAwsRte53ZoneInfo("AWS_RTE53_ZONE_INFO_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "aws_rte53_zone_info", "AWS_RTE53_ZONE_INFO_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthAwsRte53ZoneInfo("AWS_RTE53_ZONE_INFO_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "aws_rte53_zone_info", "AWS_RTE53_ZONE_INFO_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_CloudInfo(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_cloud_info"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthCloudInfo("CLOUD_INFO_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cloud_info", "CLOUD_INFO_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthCloudInfo("CLOUD_INFO_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cloud_info", "CLOUD_INFO_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_Comment(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_comment"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthComment("COMMENT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthComment("COMMENT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "COMMENT_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_CopyXferToNotify(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_copy_xfer_to_notify"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthCopyXferToNotify("COPY_XFER_TO_NOTIFY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "copy_xfer_to_notify", "COPY_XFER_TO_NOTIFY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthCopyXferToNotify("COPY_XFER_TO_NOTIFY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "copy_xfer_to_notify", "COPY_XFER_TO_NOTIFY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_CreatePtrForBulkHosts(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_create_ptr_for_bulk_hosts"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthCreatePtrForBulkHosts("CREATE_PTR_FOR_BULK_HOSTS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "create_ptr_for_bulk_hosts", "CREATE_PTR_FOR_BULK_HOSTS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthCreatePtrForBulkHosts("CREATE_PTR_FOR_BULK_HOSTS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "create_ptr_for_bulk_hosts", "CREATE_PTR_FOR_BULK_HOSTS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_CreatePtrForHosts(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_create_ptr_for_hosts"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthCreatePtrForHosts("CREATE_PTR_FOR_HOSTS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "create_ptr_for_hosts", "CREATE_PTR_FOR_HOSTS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthCreatePtrForHosts("CREATE_PTR_FOR_HOSTS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "create_ptr_for_hosts", "CREATE_PTR_FOR_HOSTS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_CreateUnderscoreZones(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_create_underscore_zones"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthCreateUnderscoreZones("CREATE_UNDERSCORE_ZONES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "create_underscore_zones", "CREATE_UNDERSCORE_ZONES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthCreateUnderscoreZones("CREATE_UNDERSCORE_ZONES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "create_underscore_zones", "CREATE_UNDERSCORE_ZONES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_DdnsForceCreationTimestampUpdate(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_ddns_force_creation_timestamp_update"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthDdnsForceCreationTimestampUpdate("DDNS_FORCE_CREATION_TIMESTAMP_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_force_creation_timestamp_update", "DDNS_FORCE_CREATION_TIMESTAMP_UPDATE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthDdnsForceCreationTimestampUpdate("DDNS_FORCE_CREATION_TIMESTAMP_UPDATE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_force_creation_timestamp_update", "DDNS_FORCE_CREATION_TIMESTAMP_UPDATE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_DdnsPrincipalGroup(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_ddns_principal_group"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthDdnsPrincipalGroup("DDNS_PRINCIPAL_GROUP_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal_group", "DDNS_PRINCIPAL_GROUP_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthDdnsPrincipalGroup("DDNS_PRINCIPAL_GROUP_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal_group", "DDNS_PRINCIPAL_GROUP_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_DdnsPrincipalTracking(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_ddns_principal_tracking"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthDdnsPrincipalTracking("DDNS_PRINCIPAL_TRACKING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal_tracking", "DDNS_PRINCIPAL_TRACKING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthDdnsPrincipalTracking("DDNS_PRINCIPAL_TRACKING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_principal_tracking", "DDNS_PRINCIPAL_TRACKING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_DdnsRestrictPatterns(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_ddns_restrict_patterns"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthDdnsRestrictPatterns("DDNS_RESTRICT_PATTERNS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_patterns", "DDNS_RESTRICT_PATTERNS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthDdnsRestrictPatterns("DDNS_RESTRICT_PATTERNS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_patterns", "DDNS_RESTRICT_PATTERNS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_DdnsRestrictPatternsList(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_ddns_restrict_patterns_list"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthDdnsRestrictPatternsList("DDNS_RESTRICT_PATTERNS_LIST_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_patterns_list", "DDNS_RESTRICT_PATTERNS_LIST_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthDdnsRestrictPatternsList("DDNS_RESTRICT_PATTERNS_LIST_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_patterns_list", "DDNS_RESTRICT_PATTERNS_LIST_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_DdnsRestrictProtected(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_ddns_restrict_protected"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthDdnsRestrictProtected("DDNS_RESTRICT_PROTECTED_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_protected", "DDNS_RESTRICT_PROTECTED_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthDdnsRestrictProtected("DDNS_RESTRICT_PROTECTED_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_protected", "DDNS_RESTRICT_PROTECTED_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_DdnsRestrictSecure(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_ddns_restrict_secure"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthDdnsRestrictSecure("DDNS_RESTRICT_SECURE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_secure", "DDNS_RESTRICT_SECURE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthDdnsRestrictSecure("DDNS_RESTRICT_SECURE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_secure", "DDNS_RESTRICT_SECURE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_DdnsRestrictStatic(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_ddns_restrict_static"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthDdnsRestrictStatic("DDNS_RESTRICT_STATIC_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_static", "DDNS_RESTRICT_STATIC_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthDdnsRestrictStatic("DDNS_RESTRICT_STATIC_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ddns_restrict_static", "DDNS_RESTRICT_STATIC_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_Disable(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_disable"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthDisable("DISABLE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthDisable("DISABLE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable", "DISABLE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_DisableForwarding(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_disable_forwarding"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthDisableForwarding("DISABLE_FORWARDING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable_forwarding", "DISABLE_FORWARDING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthDisableForwarding("DISABLE_FORWARDING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "disable_forwarding", "DISABLE_FORWARDING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_DnsIntegrityEnable(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_dns_integrity_enable"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthDnsIntegrityEnable("DNS_INTEGRITY_ENABLE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns_integrity_enable", "DNS_INTEGRITY_ENABLE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthDnsIntegrityEnable("DNS_INTEGRITY_ENABLE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns_integrity_enable", "DNS_INTEGRITY_ENABLE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_DnsIntegrityFrequency(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_dns_integrity_frequency"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthDnsIntegrityFrequency("DNS_INTEGRITY_FREQUENCY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns_integrity_frequency", "DNS_INTEGRITY_FREQUENCY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthDnsIntegrityFrequency("DNS_INTEGRITY_FREQUENCY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns_integrity_frequency", "DNS_INTEGRITY_FREQUENCY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_DnsIntegrityMember(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_dns_integrity_member"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthDnsIntegrityMember("DNS_INTEGRITY_MEMBER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns_integrity_member", "DNS_INTEGRITY_MEMBER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthDnsIntegrityMember("DNS_INTEGRITY_MEMBER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns_integrity_member", "DNS_INTEGRITY_MEMBER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_DnsIntegrityVerboseLogging(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_dns_integrity_verbose_logging"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthDnsIntegrityVerboseLogging("DNS_INTEGRITY_VERBOSE_LOGGING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns_integrity_verbose_logging", "DNS_INTEGRITY_VERBOSE_LOGGING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthDnsIntegrityVerboseLogging("DNS_INTEGRITY_VERBOSE_LOGGING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns_integrity_verbose_logging", "DNS_INTEGRITY_VERBOSE_LOGGING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_DnssecKeyParams(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_dnssec_key_params"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthDnssecKeyParams("DNSSEC_KEY_PARAMS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dnssec_key_params", "DNSSEC_KEY_PARAMS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthDnssecKeyParams("DNSSEC_KEY_PARAMS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dnssec_key_params", "DNSSEC_KEY_PARAMS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_DnssecKeys(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_dnssec_keys"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthDnssecKeys("DNSSEC_KEYS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dnssec_keys", "DNSSEC_KEYS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthDnssecKeys("DNSSEC_KEYS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dnssec_keys", "DNSSEC_KEYS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_DoHostAbstraction(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_do_host_abstraction"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthDoHostAbstraction("DO_HOST_ABSTRACTION_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "do_host_abstraction", "DO_HOST_ABSTRACTION_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthDoHostAbstraction("DO_HOST_ABSTRACTION_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "do_host_abstraction", "DO_HOST_ABSTRACTION_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_EffectiveCheckNamesPolicy(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_effective_check_names_policy"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthEffectiveCheckNamesPolicy("EFFECTIVE_CHECK_NAMES_POLICY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "effective_check_names_policy", "EFFECTIVE_CHECK_NAMES_POLICY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthEffectiveCheckNamesPolicy("EFFECTIVE_CHECK_NAMES_POLICY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "effective_check_names_policy", "EFFECTIVE_CHECK_NAMES_POLICY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_ExtAttrs(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_extattrs"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthExtAttrs("EXT_ATTRS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthExtAttrs("EXT_ATTRS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "extattrs", "EXT_ATTRS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_ExternalPrimaries(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_external_primaries"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthExternalPrimaries("EXTERNAL_PRIMARIES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_primaries", "EXTERNAL_PRIMARIES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthExternalPrimaries("EXTERNAL_PRIMARIES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_primaries", "EXTERNAL_PRIMARIES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_ExternalSecondaries(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_external_secondaries"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthExternalSecondaries("EXTERNAL_SECONDARIES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_secondaries", "EXTERNAL_SECONDARIES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthExternalSecondaries("EXTERNAL_SECONDARIES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "external_secondaries", "EXTERNAL_SECONDARIES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_Fqdn(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_fqdn"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthFqdn("FQDN_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fqdn", "FQDN_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthFqdn("FQDN_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fqdn", "FQDN_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_GridPrimary(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_grid_primary"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthGridPrimary("GRID_PRIMARY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "grid_primary", "GRID_PRIMARY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthGridPrimary("GRID_PRIMARY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "grid_primary", "GRID_PRIMARY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_GridSecondaries(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_grid_secondaries"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthGridSecondaries("GRID_SECONDARIES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "grid_secondaries", "GRID_SECONDARIES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthGridSecondaries("GRID_SECONDARIES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "grid_secondaries", "GRID_SECONDARIES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_ImportFrom(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_import_from"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthImportFrom("IMPORT_FROM_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "import_from", "IMPORT_FROM_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthImportFrom("IMPORT_FROM_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "import_from", "IMPORT_FROM_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_LastQueriedAcl(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_last_queried_acl"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthLastQueriedAcl("LAST_QUERIED_ACL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "last_queried_acl", "LAST_QUERIED_ACL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthLastQueriedAcl("LAST_QUERIED_ACL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "last_queried_acl", "LAST_QUERIED_ACL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_Locked(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_locked"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthLocked("LOCKED_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "locked", "LOCKED_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthLocked("LOCKED_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "locked", "LOCKED_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_MemberSoaMnames(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_member_soa_mnames"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthMemberSoaMnames("MEMBER_SOA_MNAMES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "member_soa_mnames", "MEMBER_SOA_MNAMES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthMemberSoaMnames("MEMBER_SOA_MNAMES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "member_soa_mnames", "MEMBER_SOA_MNAMES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_MsAdIntegrated(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_ms_ad_integrated"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthMsAdIntegrated("MS_AD_INTEGRATED_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_ad_integrated", "MS_AD_INTEGRATED_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthMsAdIntegrated("MS_AD_INTEGRATED_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_ad_integrated", "MS_AD_INTEGRATED_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_MsAllowTransfer(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_ms_allow_transfer"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthMsAllowTransfer("MS_ALLOW_TRANSFER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_allow_transfer", "MS_ALLOW_TRANSFER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthMsAllowTransfer("MS_ALLOW_TRANSFER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_allow_transfer", "MS_ALLOW_TRANSFER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_MsAllowTransferMode(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_ms_allow_transfer_mode"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthMsAllowTransferMode("MS_ALLOW_TRANSFER_MODE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_allow_transfer_mode", "MS_ALLOW_TRANSFER_MODE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthMsAllowTransferMode("MS_ALLOW_TRANSFER_MODE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_allow_transfer_mode", "MS_ALLOW_TRANSFER_MODE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_MsDcNsRecordCreation(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_ms_dc_ns_record_creation"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthMsDcNsRecordCreation("MS_DC_NS_RECORD_CREATION_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_dc_ns_record_creation", "MS_DC_NS_RECORD_CREATION_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthMsDcNsRecordCreation("MS_DC_NS_RECORD_CREATION_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_dc_ns_record_creation", "MS_DC_NS_RECORD_CREATION_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_MsDdnsMode(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_ms_ddns_mode"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthMsDdnsMode("MS_DDNS_MODE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_ddns_mode", "MS_DDNS_MODE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthMsDdnsMode("MS_DDNS_MODE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_ddns_mode", "MS_DDNS_MODE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_MsPrimaries(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_ms_primaries"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthMsPrimaries("MS_PRIMARIES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_primaries", "MS_PRIMARIES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthMsPrimaries("MS_PRIMARIES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_primaries", "MS_PRIMARIES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_MsSecondaries(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_ms_secondaries"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthMsSecondaries("MS_SECONDARIES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_secondaries", "MS_SECONDARIES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthMsSecondaries("MS_SECONDARIES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_secondaries", "MS_SECONDARIES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_MsSyncDisabled(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_ms_sync_disabled"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthMsSyncDisabled("MS_SYNC_DISABLED_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_sync_disabled", "MS_SYNC_DISABLED_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthMsSyncDisabled("MS_SYNC_DISABLED_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ms_sync_disabled", "MS_SYNC_DISABLED_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_NotifyDelay(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_notify_delay"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthNotifyDelay("NOTIFY_DELAY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "notify_delay", "NOTIFY_DELAY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthNotifyDelay("NOTIFY_DELAY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "notify_delay", "NOTIFY_DELAY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_NsGroup(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_ns_group"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthNsGroup("NS_GROUP_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ns_group", "NS_GROUP_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthNsGroup("NS_GROUP_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ns_group", "NS_GROUP_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_Prefix(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_prefix"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthPrefix("PREFIX_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "prefix", "PREFIX_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthPrefix("PREFIX_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "prefix", "PREFIX_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_RecordNamePolicy(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_record_name_policy"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthRecordNamePolicy("RECORD_NAME_POLICY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "record_name_policy", "RECORD_NAME_POLICY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthRecordNamePolicy("RECORD_NAME_POLICY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "record_name_policy", "RECORD_NAME_POLICY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_RemoveSubzones(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_remove_subzones"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthRemoveSubzones("REMOVE_SUBZONES_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "remove_subzones", "REMOVE_SUBZONES_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthRemoveSubzones("REMOVE_SUBZONES_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "remove_subzones", "REMOVE_SUBZONES_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_RestartIfNeeded(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_restart_if_needed"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthRestartIfNeeded("RESTART_IF_NEEDED_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "restart_if_needed", "RESTART_IF_NEEDED_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthRestartIfNeeded("RESTART_IF_NEEDED_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "restart_if_needed", "RESTART_IF_NEEDED_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_ScavengingSettings(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_scavenging_settings"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthScavengingSettings("SCAVENGING_SETTINGS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings", "SCAVENGING_SETTINGS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthScavengingSettings("SCAVENGING_SETTINGS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "scavenging_settings", "SCAVENGING_SETTINGS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_SetSoaSerialNumber(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_set_soa_serial_number"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthSetSoaSerialNumber("SET_SOA_SERIAL_NUMBER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "set_soa_serial_number", "SET_SOA_SERIAL_NUMBER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthSetSoaSerialNumber("SET_SOA_SERIAL_NUMBER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "set_soa_serial_number", "SET_SOA_SERIAL_NUMBER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_SoaDefaultTtl(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_soa_default_ttl"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthSoaDefaultTtl("SOA_DEFAULT_TTL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_default_ttl", "SOA_DEFAULT_TTL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthSoaDefaultTtl("SOA_DEFAULT_TTL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_default_ttl", "SOA_DEFAULT_TTL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_SoaEmail(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_soa_email"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthSoaEmail("SOA_EMAIL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_email", "SOA_EMAIL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthSoaEmail("SOA_EMAIL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_email", "SOA_EMAIL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_SoaExpire(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_soa_expire"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthSoaExpire("SOA_EXPIRE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_expire", "SOA_EXPIRE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthSoaExpire("SOA_EXPIRE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_expire", "SOA_EXPIRE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_SoaNegativeTtl(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_soa_negative_ttl"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthSoaNegativeTtl("SOA_NEGATIVE_TTL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_negative_ttl", "SOA_NEGATIVE_TTL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthSoaNegativeTtl("SOA_NEGATIVE_TTL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_negative_ttl", "SOA_NEGATIVE_TTL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_SoaRefresh(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_soa_refresh"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthSoaRefresh("SOA_REFRESH_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_refresh", "SOA_REFRESH_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthSoaRefresh("SOA_REFRESH_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_refresh", "SOA_REFRESH_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_SoaRetry(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_soa_retry"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthSoaRetry("SOA_RETRY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_retry", "SOA_RETRY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthSoaRetry("SOA_RETRY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_retry", "SOA_RETRY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_SoaSerialNumber(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_soa_serial_number"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthSoaSerialNumber("SOA_SERIAL_NUMBER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_serial_number", "SOA_SERIAL_NUMBER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthSoaSerialNumber("SOA_SERIAL_NUMBER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "soa_serial_number", "SOA_SERIAL_NUMBER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_Srgs(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_srgs"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthSrgs("SRGS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "srgs", "SRGS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthSrgs("SRGS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "srgs", "SRGS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_UpdateForwarding(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_update_forwarding"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthUpdateForwarding("UPDATE_FORWARDING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_forwarding", "UPDATE_FORWARDING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthUpdateForwarding("UPDATE_FORWARDING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_forwarding", "UPDATE_FORWARDING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_UseAllowActiveDir(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_use_allow_active_dir"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthUseAllowActiveDir("USE_ALLOW_ACTIVE_DIR_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_allow_active_dir", "USE_ALLOW_ACTIVE_DIR_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthUseAllowActiveDir("USE_ALLOW_ACTIVE_DIR_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_allow_active_dir", "USE_ALLOW_ACTIVE_DIR_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_UseAllowQuery(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_use_allow_query"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthUseAllowQuery("USE_ALLOW_QUERY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_allow_query", "USE_ALLOW_QUERY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthUseAllowQuery("USE_ALLOW_QUERY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_allow_query", "USE_ALLOW_QUERY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_UseAllowTransfer(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_use_allow_transfer"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthUseAllowTransfer("USE_ALLOW_TRANSFER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_allow_transfer", "USE_ALLOW_TRANSFER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthUseAllowTransfer("USE_ALLOW_TRANSFER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_allow_transfer", "USE_ALLOW_TRANSFER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_UseAllowUpdate(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_use_allow_update"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthUseAllowUpdate("USE_ALLOW_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_allow_update", "USE_ALLOW_UPDATE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthUseAllowUpdate("USE_ALLOW_UPDATE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_allow_update", "USE_ALLOW_UPDATE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_UseAllowUpdateForwarding(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_use_allow_update_forwarding"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthUseAllowUpdateForwarding("USE_ALLOW_UPDATE_FORWARDING_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_allow_update_forwarding", "USE_ALLOW_UPDATE_FORWARDING_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthUseAllowUpdateForwarding("USE_ALLOW_UPDATE_FORWARDING_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_allow_update_forwarding", "USE_ALLOW_UPDATE_FORWARDING_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_UseCheckNamesPolicy(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_use_check_names_policy"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthUseCheckNamesPolicy("USE_CHECK_NAMES_POLICY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_check_names_policy", "USE_CHECK_NAMES_POLICY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthUseCheckNamesPolicy("USE_CHECK_NAMES_POLICY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_check_names_policy", "USE_CHECK_NAMES_POLICY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_UseCopyXferToNotify(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_use_copy_xfer_to_notify"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthUseCopyXferToNotify("USE_COPY_XFER_TO_NOTIFY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_copy_xfer_to_notify", "USE_COPY_XFER_TO_NOTIFY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthUseCopyXferToNotify("USE_COPY_XFER_TO_NOTIFY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_copy_xfer_to_notify", "USE_COPY_XFER_TO_NOTIFY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_UseDdnsForceCreationTimestampUpdate(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_use_ddns_force_creation_timestamp_update"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthUseDdnsForceCreationTimestampUpdate("USE_DDNS_FORCE_CREATION_TIMESTAMP_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_force_creation_timestamp_update", "USE_DDNS_FORCE_CREATION_TIMESTAMP_UPDATE_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthUseDdnsForceCreationTimestampUpdate("USE_DDNS_FORCE_CREATION_TIMESTAMP_UPDATE_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_force_creation_timestamp_update", "USE_DDNS_FORCE_CREATION_TIMESTAMP_UPDATE_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_UseDdnsPatternsRestriction(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_use_ddns_patterns_restriction"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthUseDdnsPatternsRestriction("USE_DDNS_PATTERNS_RESTRICTION_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_patterns_restriction", "USE_DDNS_PATTERNS_RESTRICTION_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthUseDdnsPatternsRestriction("USE_DDNS_PATTERNS_RESTRICTION_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_patterns_restriction", "USE_DDNS_PATTERNS_RESTRICTION_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_UseDdnsPrincipalSecurity(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_use_ddns_principal_security"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthUseDdnsPrincipalSecurity("USE_DDNS_PRINCIPAL_SECURITY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_principal_security", "USE_DDNS_PRINCIPAL_SECURITY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthUseDdnsPrincipalSecurity("USE_DDNS_PRINCIPAL_SECURITY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_principal_security", "USE_DDNS_PRINCIPAL_SECURITY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_UseDdnsRestrictProtected(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_use_ddns_restrict_protected"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthUseDdnsRestrictProtected("USE_DDNS_RESTRICT_PROTECTED_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_restrict_protected", "USE_DDNS_RESTRICT_PROTECTED_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthUseDdnsRestrictProtected("USE_DDNS_RESTRICT_PROTECTED_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_restrict_protected", "USE_DDNS_RESTRICT_PROTECTED_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_UseDdnsRestrictStatic(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_use_ddns_restrict_static"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthUseDdnsRestrictStatic("USE_DDNS_RESTRICT_STATIC_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_restrict_static", "USE_DDNS_RESTRICT_STATIC_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthUseDdnsRestrictStatic("USE_DDNS_RESTRICT_STATIC_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_ddns_restrict_static", "USE_DDNS_RESTRICT_STATIC_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_UseDnssecKeyParams(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_use_dnssec_key_params"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthUseDnssecKeyParams("USE_DNSSEC_KEY_PARAMS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_dnssec_key_params", "USE_DNSSEC_KEY_PARAMS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthUseDnssecKeyParams("USE_DNSSEC_KEY_PARAMS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_dnssec_key_params", "USE_DNSSEC_KEY_PARAMS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_UseExternalPrimary(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_use_external_primary"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthUseExternalPrimary("USE_EXTERNAL_PRIMARY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_external_primary", "USE_EXTERNAL_PRIMARY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthUseExternalPrimary("USE_EXTERNAL_PRIMARY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_external_primary", "USE_EXTERNAL_PRIMARY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_UseGridZoneTimer(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_use_grid_zone_timer"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthUseGridZoneTimer("USE_GRID_ZONE_TIMER_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_grid_zone_timer", "USE_GRID_ZONE_TIMER_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthUseGridZoneTimer("USE_GRID_ZONE_TIMER_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_grid_zone_timer", "USE_GRID_ZONE_TIMER_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_UseImportFrom(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_use_import_from"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthUseImportFrom("USE_IMPORT_FROM_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_import_from", "USE_IMPORT_FROM_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthUseImportFrom("USE_IMPORT_FROM_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_import_from", "USE_IMPORT_FROM_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_UseNotifyDelay(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_use_notify_delay"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthUseNotifyDelay("USE_NOTIFY_DELAY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_notify_delay", "USE_NOTIFY_DELAY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthUseNotifyDelay("USE_NOTIFY_DELAY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_notify_delay", "USE_NOTIFY_DELAY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_UseRecordNamePolicy(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_use_record_name_policy"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthUseRecordNamePolicy("USE_RECORD_NAME_POLICY_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_record_name_policy", "USE_RECORD_NAME_POLICY_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthUseRecordNamePolicy("USE_RECORD_NAME_POLICY_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_record_name_policy", "USE_RECORD_NAME_POLICY_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_UseScavengingSettings(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_use_scavenging_settings"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthUseScavengingSettings("USE_SCAVENGING_SETTINGS_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_scavenging_settings", "USE_SCAVENGING_SETTINGS_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthUseScavengingSettings("USE_SCAVENGING_SETTINGS_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_scavenging_settings", "USE_SCAVENGING_SETTINGS_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_UseSoaEmail(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_use_soa_email"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthUseSoaEmail("USE_SOA_EMAIL_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_soa_email", "USE_SOA_EMAIL_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthUseSoaEmail("USE_SOA_EMAIL_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_soa_email", "USE_SOA_EMAIL_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_View(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_view"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthView("VIEW_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", "VIEW_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthView("VIEW_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "view", "VIEW_UPDATE_REPLACE_ME"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccZoneAuthResource_ZoneFormat(t *testing.T) {
	var resourceName = "nios_dns_zone_auth.test_zone_format"
	var v dns.ZoneAuth

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccZoneAuthZoneFormat("ZONE_FORMAT_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "zone_format", "ZONE_FORMAT_REPLACE_ME"),
				),
			},
			// Update and Read
			{
				Config: testAccZoneAuthZoneFormat("ZONE_FORMAT_UPDATE_REPLACE_ME"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckZoneAuthExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "zone_format", "ZONE_FORMAT_UPDATE_REPLACE_ME"),
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

func testAccZoneAuthRef(ref string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test__ref" {
    _ref = %q
}
`, ref)
}

func testAccZoneAuthAllowActiveDir(allowActiveDir string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_allow_active_dir" {
    allow_active_dir = %q
}
`, allowActiveDir)
}

func testAccZoneAuthAllowFixedRrsetOrder(allowFixedRrsetOrder string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_allow_fixed_rrset_order" {
    allow_fixed_rrset_order = %q
}
`, allowFixedRrsetOrder)
}

func testAccZoneAuthAllowGssTsigForUnderscoreZone(allowGssTsigForUnderscoreZone string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_allow_gss_tsig_for_underscore_zone" {
    allow_gss_tsig_for_underscore_zone = %q
}
`, allowGssTsigForUnderscoreZone)
}

func testAccZoneAuthAllowGssTsigZoneUpdates(allowGssTsigZoneUpdates string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_allow_gss_tsig_zone_updates" {
    allow_gss_tsig_zone_updates = %q
}
`, allowGssTsigZoneUpdates)
}

func testAccZoneAuthAllowQuery(allowQuery string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_allow_query" {
    allow_query = %q
}
`, allowQuery)
}

func testAccZoneAuthAllowTransfer(allowTransfer string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_allow_transfer" {
    allow_transfer = %q
}
`, allowTransfer)
}

func testAccZoneAuthAllowUpdate(allowUpdate string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_allow_update" {
    allow_update = %q
}
`, allowUpdate)
}

func testAccZoneAuthAllowUpdateForwarding(allowUpdateForwarding string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_allow_update_forwarding" {
    allow_update_forwarding = %q
}
`, allowUpdateForwarding)
}

func testAccZoneAuthAwsRte53ZoneInfo(awsRte53ZoneInfo string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_aws_rte53_zone_info" {
    aws_rte53_zone_info = %q
}
`, awsRte53ZoneInfo)
}

func testAccZoneAuthCloudInfo(cloudInfo string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_cloud_info" {
    cloud_info = %q
}
`, cloudInfo)
}

func testAccZoneAuthComment(comment string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_comment" {
    comment = %q
}
`, comment)
}

func testAccZoneAuthCopyXferToNotify(copyXferToNotify string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_copy_xfer_to_notify" {
    copy_xfer_to_notify = %q
}
`, copyXferToNotify)
}

func testAccZoneAuthCreatePtrForBulkHosts(createPtrForBulkHosts string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_create_ptr_for_bulk_hosts" {
    create_ptr_for_bulk_hosts = %q
}
`, createPtrForBulkHosts)
}

func testAccZoneAuthCreatePtrForHosts(createPtrForHosts string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_create_ptr_for_hosts" {
    create_ptr_for_hosts = %q
}
`, createPtrForHosts)
}

func testAccZoneAuthCreateUnderscoreZones(createUnderscoreZones string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_create_underscore_zones" {
    create_underscore_zones = %q
}
`, createUnderscoreZones)
}

func testAccZoneAuthDdnsForceCreationTimestampUpdate(ddnsForceCreationTimestampUpdate string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_ddns_force_creation_timestamp_update" {
    ddns_force_creation_timestamp_update = %q
}
`, ddnsForceCreationTimestampUpdate)
}

func testAccZoneAuthDdnsPrincipalGroup(ddnsPrincipalGroup string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_ddns_principal_group" {
    ddns_principal_group = %q
}
`, ddnsPrincipalGroup)
}

func testAccZoneAuthDdnsPrincipalTracking(ddnsPrincipalTracking string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_ddns_principal_tracking" {
    ddns_principal_tracking = %q
}
`, ddnsPrincipalTracking)
}

func testAccZoneAuthDdnsRestrictPatterns(ddnsRestrictPatterns string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_ddns_restrict_patterns" {
    ddns_restrict_patterns = %q
}
`, ddnsRestrictPatterns)
}

func testAccZoneAuthDdnsRestrictPatternsList(ddnsRestrictPatternsList string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_ddns_restrict_patterns_list" {
    ddns_restrict_patterns_list = %q
}
`, ddnsRestrictPatternsList)
}

func testAccZoneAuthDdnsRestrictProtected(ddnsRestrictProtected string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_ddns_restrict_protected" {
    ddns_restrict_protected = %q
}
`, ddnsRestrictProtected)
}

func testAccZoneAuthDdnsRestrictSecure(ddnsRestrictSecure string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_ddns_restrict_secure" {
    ddns_restrict_secure = %q
}
`, ddnsRestrictSecure)
}

func testAccZoneAuthDdnsRestrictStatic(ddnsRestrictStatic string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_ddns_restrict_static" {
    ddns_restrict_static = %q
}
`, ddnsRestrictStatic)
}

func testAccZoneAuthDisable(disable string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_disable" {
    disable = %q
}
`, disable)
}

func testAccZoneAuthDisableForwarding(disableForwarding string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_disable_forwarding" {
    disable_forwarding = %q
}
`, disableForwarding)
}

func testAccZoneAuthDnsIntegrityEnable(dnsIntegrityEnable string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_dns_integrity_enable" {
    dns_integrity_enable = %q
}
`, dnsIntegrityEnable)
}

func testAccZoneAuthDnsIntegrityFrequency(dnsIntegrityFrequency string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_dns_integrity_frequency" {
    dns_integrity_frequency = %q
}
`, dnsIntegrityFrequency)
}

func testAccZoneAuthDnsIntegrityMember(dnsIntegrityMember string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_dns_integrity_member" {
    dns_integrity_member = %q
}
`, dnsIntegrityMember)
}

func testAccZoneAuthDnsIntegrityVerboseLogging(dnsIntegrityVerboseLogging string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_dns_integrity_verbose_logging" {
    dns_integrity_verbose_logging = %q
}
`, dnsIntegrityVerboseLogging)
}

func testAccZoneAuthDnssecKeyParams(dnssecKeyParams string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_dnssec_key_params" {
    dnssec_key_params = %q
}
`, dnssecKeyParams)
}

func testAccZoneAuthDnssecKeys(dnssecKeys string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_dnssec_keys" {
    dnssec_keys = %q
}
`, dnssecKeys)
}

func testAccZoneAuthDoHostAbstraction(doHostAbstraction string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_do_host_abstraction" {
    do_host_abstraction = %q
}
`, doHostAbstraction)
}

func testAccZoneAuthEffectiveCheckNamesPolicy(effectiveCheckNamesPolicy string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_effective_check_names_policy" {
    effective_check_names_policy = %q
}
`, effectiveCheckNamesPolicy)
}

func testAccZoneAuthExtAttrs(extAttrs string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_extattrs" {
    extattrs = %q
}
`, extAttrs)
}

func testAccZoneAuthExternalPrimaries(externalPrimaries string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_external_primaries" {
    external_primaries = %q
}
`, externalPrimaries)
}

func testAccZoneAuthExternalSecondaries(externalSecondaries string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_external_secondaries" {
    external_secondaries = %q
}
`, externalSecondaries)
}

func testAccZoneAuthFqdn(fqdn string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_fqdn" {
    fqdn = %q
}
`, fqdn)
}

func testAccZoneAuthGridPrimary(gridPrimary string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_grid_primary" {
    grid_primary = %q
}
`, gridPrimary)
}

func testAccZoneAuthGridSecondaries(gridSecondaries string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_grid_secondaries" {
    grid_secondaries = %q
}
`, gridSecondaries)
}

func testAccZoneAuthImportFrom(importFrom string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_import_from" {
    import_from = %q
}
`, importFrom)
}

func testAccZoneAuthLastQueriedAcl(lastQueriedAcl string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_last_queried_acl" {
    last_queried_acl = %q
}
`, lastQueriedAcl)
}

func testAccZoneAuthLocked(locked string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_locked" {
    locked = %q
}
`, locked)
}

func testAccZoneAuthMemberSoaMnames(memberSoaMnames string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_member_soa_mnames" {
    member_soa_mnames = %q
}
`, memberSoaMnames)
}

func testAccZoneAuthMsAdIntegrated(msAdIntegrated string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_ms_ad_integrated" {
    ms_ad_integrated = %q
}
`, msAdIntegrated)
}

func testAccZoneAuthMsAllowTransfer(msAllowTransfer string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_ms_allow_transfer" {
    ms_allow_transfer = %q
}
`, msAllowTransfer)
}

func testAccZoneAuthMsAllowTransferMode(msAllowTransferMode string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_ms_allow_transfer_mode" {
    ms_allow_transfer_mode = %q
}
`, msAllowTransferMode)
}

func testAccZoneAuthMsDcNsRecordCreation(msDcNsRecordCreation string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_ms_dc_ns_record_creation" {
    ms_dc_ns_record_creation = %q
}
`, msDcNsRecordCreation)
}

func testAccZoneAuthMsDdnsMode(msDdnsMode string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_ms_ddns_mode" {
    ms_ddns_mode = %q
}
`, msDdnsMode)
}

func testAccZoneAuthMsPrimaries(msPrimaries string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_ms_primaries" {
    ms_primaries = %q
}
`, msPrimaries)
}

func testAccZoneAuthMsSecondaries(msSecondaries string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_ms_secondaries" {
    ms_secondaries = %q
}
`, msSecondaries)
}

func testAccZoneAuthMsSyncDisabled(msSyncDisabled string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_ms_sync_disabled" {
    ms_sync_disabled = %q
}
`, msSyncDisabled)
}

func testAccZoneAuthNotifyDelay(notifyDelay string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_notify_delay" {
    notify_delay = %q
}
`, notifyDelay)
}

func testAccZoneAuthNsGroup(nsGroup string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_ns_group" {
    ns_group = %q
}
`, nsGroup)
}

func testAccZoneAuthPrefix(prefix string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_prefix" {
    prefix = %q
}
`, prefix)
}

func testAccZoneAuthRecordNamePolicy(recordNamePolicy string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_record_name_policy" {
    record_name_policy = %q
}
`, recordNamePolicy)
}

func testAccZoneAuthRemoveSubzones(removeSubzones string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_remove_subzones" {
    remove_subzones = %q
}
`, removeSubzones)
}

func testAccZoneAuthRestartIfNeeded(restartIfNeeded string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_restart_if_needed" {
    restart_if_needed = %q
}
`, restartIfNeeded)
}

func testAccZoneAuthScavengingSettings(scavengingSettings string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_scavenging_settings" {
    scavenging_settings = %q
}
`, scavengingSettings)
}

func testAccZoneAuthSetSoaSerialNumber(setSoaSerialNumber string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_set_soa_serial_number" {
    set_soa_serial_number = %q
}
`, setSoaSerialNumber)
}

func testAccZoneAuthSoaDefaultTtl(soaDefaultTtl string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_soa_default_ttl" {
    soa_default_ttl = %q
}
`, soaDefaultTtl)
}

func testAccZoneAuthSoaEmail(soaEmail string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_soa_email" {
    soa_email = %q
}
`, soaEmail)
}

func testAccZoneAuthSoaExpire(soaExpire string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_soa_expire" {
    soa_expire = %q
}
`, soaExpire)
}

func testAccZoneAuthSoaNegativeTtl(soaNegativeTtl string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_soa_negative_ttl" {
    soa_negative_ttl = %q
}
`, soaNegativeTtl)
}

func testAccZoneAuthSoaRefresh(soaRefresh string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_soa_refresh" {
    soa_refresh = %q
}
`, soaRefresh)
}

func testAccZoneAuthSoaRetry(soaRetry string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_soa_retry" {
    soa_retry = %q
}
`, soaRetry)
}

func testAccZoneAuthSoaSerialNumber(soaSerialNumber string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_soa_serial_number" {
    soa_serial_number = %q
}
`, soaSerialNumber)
}

func testAccZoneAuthSrgs(srgs string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_srgs" {
    srgs = %q
}
`, srgs)
}

func testAccZoneAuthUpdateForwarding(updateForwarding string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_update_forwarding" {
    update_forwarding = %q
}
`, updateForwarding)
}

func testAccZoneAuthUseAllowActiveDir(useAllowActiveDir string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_use_allow_active_dir" {
    use_allow_active_dir = %q
}
`, useAllowActiveDir)
}

func testAccZoneAuthUseAllowQuery(useAllowQuery string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_use_allow_query" {
    use_allow_query = %q
}
`, useAllowQuery)
}

func testAccZoneAuthUseAllowTransfer(useAllowTransfer string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_use_allow_transfer" {
    use_allow_transfer = %q
}
`, useAllowTransfer)
}

func testAccZoneAuthUseAllowUpdate(useAllowUpdate string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_use_allow_update" {
    use_allow_update = %q
}
`, useAllowUpdate)
}

func testAccZoneAuthUseAllowUpdateForwarding(useAllowUpdateForwarding string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_use_allow_update_forwarding" {
    use_allow_update_forwarding = %q
}
`, useAllowUpdateForwarding)
}

func testAccZoneAuthUseCheckNamesPolicy(useCheckNamesPolicy string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_use_check_names_policy" {
    use_check_names_policy = %q
}
`, useCheckNamesPolicy)
}

func testAccZoneAuthUseCopyXferToNotify(useCopyXferToNotify string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_use_copy_xfer_to_notify" {
    use_copy_xfer_to_notify = %q
}
`, useCopyXferToNotify)
}

func testAccZoneAuthUseDdnsForceCreationTimestampUpdate(useDdnsForceCreationTimestampUpdate string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_use_ddns_force_creation_timestamp_update" {
    use_ddns_force_creation_timestamp_update = %q
}
`, useDdnsForceCreationTimestampUpdate)
}

func testAccZoneAuthUseDdnsPatternsRestriction(useDdnsPatternsRestriction string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_use_ddns_patterns_restriction" {
    use_ddns_patterns_restriction = %q
}
`, useDdnsPatternsRestriction)
}

func testAccZoneAuthUseDdnsPrincipalSecurity(useDdnsPrincipalSecurity string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_use_ddns_principal_security" {
    use_ddns_principal_security = %q
}
`, useDdnsPrincipalSecurity)
}

func testAccZoneAuthUseDdnsRestrictProtected(useDdnsRestrictProtected string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_use_ddns_restrict_protected" {
    use_ddns_restrict_protected = %q
}
`, useDdnsRestrictProtected)
}

func testAccZoneAuthUseDdnsRestrictStatic(useDdnsRestrictStatic string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_use_ddns_restrict_static" {
    use_ddns_restrict_static = %q
}
`, useDdnsRestrictStatic)
}

func testAccZoneAuthUseDnssecKeyParams(useDnssecKeyParams string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_use_dnssec_key_params" {
    use_dnssec_key_params = %q
}
`, useDnssecKeyParams)
}

func testAccZoneAuthUseExternalPrimary(useExternalPrimary string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_use_external_primary" {
    use_external_primary = %q
}
`, useExternalPrimary)
}

func testAccZoneAuthUseGridZoneTimer(useGridZoneTimer string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_use_grid_zone_timer" {
    use_grid_zone_timer = %q
}
`, useGridZoneTimer)
}

func testAccZoneAuthUseImportFrom(useImportFrom string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_use_import_from" {
    use_import_from = %q
}
`, useImportFrom)
}

func testAccZoneAuthUseNotifyDelay(useNotifyDelay string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_use_notify_delay" {
    use_notify_delay = %q
}
`, useNotifyDelay)
}

func testAccZoneAuthUseRecordNamePolicy(useRecordNamePolicy string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_use_record_name_policy" {
    use_record_name_policy = %q
}
`, useRecordNamePolicy)
}

func testAccZoneAuthUseScavengingSettings(useScavengingSettings string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_use_scavenging_settings" {
    use_scavenging_settings = %q
}
`, useScavengingSettings)
}

func testAccZoneAuthUseSoaEmail(useSoaEmail string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_use_soa_email" {
    use_soa_email = %q
}
`, useSoaEmail)
}

func testAccZoneAuthView(view string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_view" {
    view = %q
}
`, view)
}

func testAccZoneAuthZoneFormat(zoneFormat string) string {
	return fmt.Sprintf(`
resource "nios_dns_zone_auth" "test_zone_format" {
    zone_format = %q
}
`, zoneFormat)
}
