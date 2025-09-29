package discovery_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/infoblox-nios-go-client/discovery"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
)

func TestAccVdiscoverytaskDataSource_Filters(t *testing.T) {
	dataSourceName := "data.nios_discovery_vdiscoverytask.test"
	resourceName := "nios_discovery_vdiscoverytask.test"
	var v discovery.Vdiscoverytask

	name := acctest.RandomNameWithPrefix("vdiscoverytask-")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVdiscoverytaskDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccVdiscoverytaskDataSourceConfigFilters(name, true, true, true, "infoblox.172_28_83_29", "AWS", "AUTO_CREATE", "AUTO_CREATE", true, false, "us-east-1", "aws_access_key", "aws_secret_key"),
				Check: resource.ComposeTestCheckFunc(
					append([]resource.TestCheckFunc{
						testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					}, testAccCheckVdiscoverytaskResourceAttrPair(resourceName, dataSourceName)...)...,
				),
			},
		},
	})
}

// below all TestAcc functions

func testAccCheckVdiscoverytaskResourceAttrPair(resourceName, dataSourceName string) []resource.TestCheckFunc {
	return []resource.TestCheckFunc{
		resource.TestCheckResourceAttrPair(resourceName, "ref", dataSourceName, "result.0.ref"),
		resource.TestCheckResourceAttrPair(resourceName, "accounts_list", dataSourceName, "result.0.accounts_list"),
		resource.TestCheckResourceAttrPair(resourceName, "allow_unsecured_connection", dataSourceName, "result.0.allow_unsecured_connection"),
		resource.TestCheckResourceAttrPair(resourceName, "auto_consolidate_cloud_ea", dataSourceName, "result.0.auto_consolidate_cloud_ea"),
		resource.TestCheckResourceAttrPair(resourceName, "auto_consolidate_managed_tenant", dataSourceName, "result.0.auto_consolidate_managed_tenant"),
		resource.TestCheckResourceAttrPair(resourceName, "auto_consolidate_managed_vm", dataSourceName, "result.0.auto_consolidate_managed_vm"),
		resource.TestCheckResourceAttrPair(resourceName, "auto_create_dns_hostname_template", dataSourceName, "result.0.auto_create_dns_hostname_template"),
		resource.TestCheckResourceAttrPair(resourceName, "auto_create_dns_record", dataSourceName, "result.0.auto_create_dns_record"),
		resource.TestCheckResourceAttrPair(resourceName, "auto_create_dns_record_type", dataSourceName, "result.0.auto_create_dns_record_type"),
		resource.TestCheckResourceAttrPair(resourceName, "cdiscovery_file_token", dataSourceName, "result.0.cdiscovery_file_token"),
		resource.TestCheckResourceAttrPair(resourceName, "comment", dataSourceName, "result.0.comment"),
		resource.TestCheckResourceAttrPair(resourceName, "credentials_type", dataSourceName, "result.0.credentials_type"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_view_private_ip", dataSourceName, "result.0.dns_view_private_ip"),
		resource.TestCheckResourceAttrPair(resourceName, "dns_view_public_ip", dataSourceName, "result.0.dns_view_public_ip"),
		resource.TestCheckResourceAttrPair(resourceName, "domain_name", dataSourceName, "result.0.domain_name"),
		resource.TestCheckResourceAttrPair(resourceName, "driver_type", dataSourceName, "result.0.driver_type"),
		resource.TestCheckResourceAttrPair(resourceName, "enable_filter", dataSourceName, "result.0.enable_filter"),
		resource.TestCheckResourceAttrPair(resourceName, "enabled", dataSourceName, "result.0.enabled"),
		resource.TestCheckResourceAttrPair(resourceName, "fqdn_or_ip", dataSourceName, "result.0.fqdn_or_ip"),
		resource.TestCheckResourceAttrPair(resourceName, "govcloud_enabled", dataSourceName, "result.0.govcloud_enabled"),
		resource.TestCheckResourceAttrPair(resourceName, "identity_version", dataSourceName, "result.0.identity_version"),
		resource.TestCheckResourceAttrPair(resourceName, "last_run", dataSourceName, "result.0.last_run"),
		resource.TestCheckResourceAttrPair(resourceName, "member", dataSourceName, "result.0.member"),
		resource.TestCheckResourceAttrPair(resourceName, "merge_data", dataSourceName, "result.0.merge_data"),
		resource.TestCheckResourceAttrPair(resourceName, "multiple_accounts_sync_policy", dataSourceName, "result.0.multiple_accounts_sync_policy"),
		resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "result.0.name"),
		resource.TestCheckResourceAttrPair(resourceName, "network_filter", dataSourceName, "result.0.network_filter"),
		resource.TestCheckResourceAttrPair(resourceName, "network_list", dataSourceName, "result.0.network_list"),
		resource.TestCheckResourceAttrPair(resourceName, "port", dataSourceName, "result.0.port"),
		resource.TestCheckResourceAttrPair(resourceName, "private_network_view", dataSourceName, "result.0.private_network_view"),
		resource.TestCheckResourceAttrPair(resourceName, "private_network_view_mapping_policy", dataSourceName, "result.0.private_network_view_mapping_policy"),
		resource.TestCheckResourceAttrPair(resourceName, "protocol", dataSourceName, "result.0.protocol"),
		resource.TestCheckResourceAttrPair(resourceName, "public_network_view", dataSourceName, "result.0.public_network_view"),
		resource.TestCheckResourceAttrPair(resourceName, "public_network_view_mapping_policy", dataSourceName, "result.0.public_network_view_mapping_policy"),
		resource.TestCheckResourceAttrPair(resourceName, "role_arn", dataSourceName, "result.0.role_arn"),
		resource.TestCheckResourceAttrPair(resourceName, "scheduled_run", dataSourceName, "result.0.scheduled_run"),
		resource.TestCheckResourceAttrPair(resourceName, "selected_regions", dataSourceName, "result.0.selected_regions"),
		resource.TestCheckResourceAttrPair(resourceName, "service_account_file", dataSourceName, "result.0.service_account_file"),
		resource.TestCheckResourceAttrPair(resourceName, "service_account_file_token", dataSourceName, "result.0.service_account_file_token"),
		resource.TestCheckResourceAttrPair(resourceName, "state", dataSourceName, "result.0.state"),
		resource.TestCheckResourceAttrPair(resourceName, "state_msg", dataSourceName, "result.0.state_msg"),
		resource.TestCheckResourceAttrPair(resourceName, "sync_child_accounts", dataSourceName, "result.0.sync_child_accounts"),
		resource.TestCheckResourceAttrPair(resourceName, "update_dns_view_private_ip", dataSourceName, "result.0.update_dns_view_private_ip"),
		resource.TestCheckResourceAttrPair(resourceName, "update_dns_view_public_ip", dataSourceName, "result.0.update_dns_view_public_ip"),
		resource.TestCheckResourceAttrPair(resourceName, "update_metadata", dataSourceName, "result.0.update_metadata"),
		resource.TestCheckResourceAttrPair(resourceName, "use_identity", dataSourceName, "result.0.use_identity"),
		resource.TestCheckResourceAttrPair(resourceName, "username", dataSourceName, "result.0.username"),
	}
}

func testAccVdiscoverytaskDataSourceConfigFilters(name string, auto_consolidate_cloud_ea, auto_consolidate_managed_tenant, auto_consolidate_managed_vm bool, member, driver_type, private_network_view_mapping_policy, public_network_view_mapping_policy string, merge_data, update_metadata bool, selected_regions, username, password string) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscoverytask" "test" {
    name                                = %q
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    driver_type                         = %q
    member                              = %q
    merge_data                          = %t
    private_network_view_mapping_policy = %q
    public_network_view_mapping_policy  = %q
    update_metadata                     = %t
    selected_regions                    = %q
    username                            = %q
	password                            = %q
}

data "nios_discovery_vdiscoverytask" "test" {
	filters = {
    	name = nios_discovery_vdiscoverytask.test.name
  	}
}
`, name, auto_consolidate_cloud_ea, auto_consolidate_managed_tenant, auto_consolidate_managed_vm, driver_type, member, merge_data, private_network_view_mapping_policy, public_network_view_mapping_policy, update_metadata, selected_regions, username, password)
}
