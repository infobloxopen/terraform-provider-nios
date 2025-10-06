package discovery_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/infoblox-nios-go-client/discovery"
	"github.com/infobloxopen/terraform-provider-nios/internal/acctest"
	"github.com/infobloxopen/terraform-provider-nios/internal/utils"
)

// TODO: OBJECTS TO BE PRESENT IN THE GRID FOR TESTS
// Network views: custom_private_view, custom_public_view
// DNS view : custom_dns_view
// GCP service account files
// CSV files for cdiscovery if muliple accounts sync policy is UPLOAD

var readableAttributesForVdiscoverytask = "accounts_list,allow_unsecured_connection,auto_consolidate_cloud_ea,auto_consolidate_managed_tenant,auto_consolidate_managed_vm,auto_create_dns_hostname_template,auto_create_dns_record,auto_create_dns_record_type,cdiscovery_file_token,comment,credentials_type,dns_view_private_ip,dns_view_public_ip,domain_name,driver_type,enable_filter,enabled,fqdn_or_ip,govcloud_enabled,identity_version,last_run,member,merge_data,multiple_accounts_sync_policy,name,network_filter,network_list,port,private_network_view,private_network_view_mapping_policy,protocol,public_network_view,public_network_view_mapping_policy,role_arn,scheduled_run,selected_regions,service_account_file,service_account_file_token,state,state_msg,sync_child_accounts,update_dns_view_private_ip,update_dns_view_public_ip,update_metadata,use_identity,username"

func TestAccVdiscoverytaskResource_basic(t *testing.T) {
	var resourceName = "nios_discovery_vdiscovery_task.test"
	var v discovery.Vdiscoverytask

	name := acctest.RandomNameWithPrefix("example-vdiscoverytask-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVdiscoverytaskBasicConfig(name, true, true, true, "AWS", "infoblox.172_28_83_29", "AUTO_CREATE", "AUTO_CREATE", true, false, "us-east-1", "aws_access_key", "aws_secret_key"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "driver_type", "AWS"),
					resource.TestCheckResourceAttr(resourceName, "member", "infoblox.172_28_83_29"),
					resource.TestCheckResourceAttr(resourceName, "auto_consolidate_cloud_ea", "true"),
					resource.TestCheckResourceAttr(resourceName, "auto_consolidate_managed_tenant", "true"),
					resource.TestCheckResourceAttr(resourceName, "auto_consolidate_managed_vm", "true"),
					resource.TestCheckResourceAttr(resourceName, "merge_data", "true"),
					resource.TestCheckResourceAttr(resourceName, "private_network_view_mapping_policy", "AUTO_CREATE"),
					resource.TestCheckResourceAttr(resourceName, "public_network_view_mapping_policy", "AUTO_CREATE"),
					resource.TestCheckResourceAttr(resourceName, "update_metadata", "false"),
					resource.TestCheckResourceAttr(resourceName, "selected_regions", "us-east-1"),
					resource.TestCheckResourceAttr(resourceName, "username", "aws_access_key"),
					// Test Default Values
					resource.TestCheckResourceAttr(resourceName, "allow_unsecured_connection", "false"),
					resource.TestCheckResourceAttr(resourceName, "enable_filter", "false"),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "govcloud_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "multiple_accounts_sync_policy", "DISCOVER"),
					resource.TestCheckResourceAttr(resourceName, "network_filter", "NONE"),
					resource.TestCheckResourceAttr(resourceName, "port", "443"),
					resource.TestCheckResourceAttr(resourceName, "protocol", "HTTPS"),
					resource.TestCheckResourceAttr(resourceName, "role_arn", ""),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVdiscoverytaskResource_disappears(t *testing.T) {
	resourceName := "nios_discovery_vdiscovery_task.test"
	var v discovery.Vdiscoverytask

	name := acctest.RandomNameWithPrefix("example-vdiscoverytask-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckVdiscoverytaskDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccVdiscoverytaskBasicConfig(name, true, true, true, "AWS", "infoblox.172_28_83_29", "AUTO_CREATE", "AUTO_CREATE", true, false, "us-east-1", "aws_access_key", "aws_secret_key"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					testAccCheckVdiscoverytaskDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccVdiscoverytaskResource_AllowUnsecuredConnection(t *testing.T) {
	var resourceName = "nios_discovery_vdiscovery_task.test_allow_unsecured_connection"
	var v discovery.Vdiscoverytask

	name := acctest.RandomNameWithPrefix("example-vdiscoverytask-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVdiscoverytaskAllowUnsecuredConnection(name, true, true, true, true, true, false, "VMWARE", "vcenter.example.com", "infoblox.172_28_83_29", "vmware_password", "AUTO_CREATE", "HTTPS", "AUTO_CREATE", "us-east-1", "vc_admin", 443),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_unsecured_connection", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccVdiscoverytaskAllowUnsecuredConnection(name, false, true, true, true, true, false, "VMWARE", "vcenter.example.com", "infoblox.172_28_83_29", "vmware_password", "AUTO_CREATE", "HTTPS", "AUTO_CREATE", "us-east-1", "vc_admin", 443),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "allow_unsecured_connection", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVdiscoverytaskResource_AutoConsolidateCloudEa(t *testing.T) {
	var resourceName = "nios_discovery_vdiscovery_task.test_auto_consolidate_cloud_ea"
	var v discovery.Vdiscoverytask

	name := acctest.RandomNameWithPrefix("example-vdiscoverytask-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVdiscoverytaskAutoConsolidateCloudEa(name, true, true, true, true, false, "AWS", "infoblox.172_28_83_29", "AUTO_CREATE", "AUTO_CREATE", "us-east-1", "aws_access_key", "aws_secret_key"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auto_consolidate_cloud_ea", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccVdiscoverytaskAutoConsolidateCloudEa(name, false, true, true, true, false, "AWS", "infoblox.172_28_83_29", "AUTO_CREATE", "AUTO_CREATE", "us-east-1", "aws_access_key", "aws_secret_key"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auto_consolidate_cloud_ea", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVdiscoverytaskResource_AutoConsolidateManagedTenant(t *testing.T) {
	var resourceName = "nios_discovery_vdiscovery_task.test_auto_consolidate_managed_tenant"
	var v discovery.Vdiscoverytask

	name := acctest.RandomNameWithPrefix("example-vdiscoverytask-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVdiscoverytaskAutoConsolidateManagedTenant(name, true, true, true, true, false, "AWS", "infoblox.172_28_83_29", "AUTO_CREATE", "AUTO_CREATE", "us-east-1", "aws_access_key", "aws_secret_key"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auto_consolidate_managed_tenant", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccVdiscoverytaskAutoConsolidateManagedTenant(name, false, true, true, true, false, "AWS", "infoblox.172_28_83_29", "AUTO_CREATE", "AUTO_CREATE", "us-east-1", "aws_access_key", "aws_secret_key"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auto_consolidate_managed_tenant", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVdiscoverytaskResource_AutoConsolidateManagedVm(t *testing.T) {
	var resourceName = "nios_discovery_vdiscovery_task.test_auto_consolidate_managed_vm"
	var v discovery.Vdiscoverytask

	name := acctest.RandomNameWithPrefix("example-vdiscoverytask-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVdiscoverytaskAutoConsolidateManagedVm(name, true, true, true, true, false, "AWS", "infoblox.172_28_83_29", "AUTO_CREATE", "AUTO_CREATE", "us-east-1", "aws_access_key", "aws_secret_key"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auto_consolidate_managed_vm", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccVdiscoverytaskAutoConsolidateManagedVm(name, false, true, true, true, false, "AWS", "infoblox.172_28_83_29", "AUTO_CREATE", "AUTO_CREATE", "us-east-1", "aws_access_key", "aws_secret_key"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auto_consolidate_managed_vm", "false"),
				),
			},
		},
	})
}

func TestAccVdiscoverytaskResource_AutoCreateDnsHostnameTemplate(t *testing.T) {
	var resourceName = "nios_discovery_vdiscovery_task.test_auto_create_dns_hostname_template"
	var v discovery.Vdiscoverytask

	name := acctest.RandomNameWithPrefix("example-vdiscoverytask-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVdiscoverytaskAutoCreateDnsHostnameTemplate(name, "$${vm_name}.testdomain.com", true, true, true, true, "HOST_RECORD", "AWS", "infoblox.172_28_83_29", "AUTO_CREATE", "AUTO_CREATE", "ap-northeast-1", "aws_access_key", "aws_secret_key", true, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auto_create_dns_hostname_template", "${vm_name}.testdomain.com"),
				),
			},
			// Update and Read
			{
				Config: testAccVdiscoverytaskAutoCreateDnsHostnameTemplate(name, "$${vm_name}.updatedtestdomain.com", true, true, true, true, "HOST_RECORD", "AWS", "infoblox.172_28_83_29", "AUTO_CREATE", "AUTO_CREATE", "ap-northeast-1", "aws_access_key", "aws_secret_key", true, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auto_create_dns_hostname_template", "${vm_name}.updatedtestdomain.com"),
				),
			},
		},
	})
}

func TestAccVdiscoverytaskResource_AutoCreateDnsRecord(t *testing.T) {
	var resourceName = "nios_discovery_vdiscovery_task.test_auto_create_dns_record"
	var v discovery.Vdiscoverytask

	name := acctest.RandomNameWithPrefix("example-vdiscoverytask-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVdiscoverytaskAutoCreateDnsRecord(name, true, true, true, true, "$${vm_name}.testdomain.com", "HOST_RECORD", "AWS", "infoblox.172_28_83_29", "AUTO_CREATE", "AUTO_CREATE", "ap-northeast-1", "aws_access_key", "aws_secret_key", true, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auto_create_dns_record", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccVdiscoverytaskAutoCreateDnsRecord(name, false, true, true, true, "$${vm_name}.testdomain.com", "HOST_RECORD", "AWS", "infoblox.172_28_83_29", "AUTO_CREATE", "AUTO_CREATE", "ap-northeast-1", "aws_access_key", "aws_secret_key", true, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auto_create_dns_record", "false"),
				),
			},
		},
	})
}

func TestAccVdiscoverytaskResource_AutoCreateDnsRecordType(t *testing.T) {
	var resourceName = "nios_discovery_vdiscovery_task.test_auto_create_dns_record_type"
	var v discovery.Vdiscoverytask

	name := acctest.RandomNameWithPrefix("example-vdiscoverytask-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVdiscoverytaskAutoCreateDnsRecordType(name, "HOST_RECORD", true, true, true, true, "$${vm_name}.testdomain.com", "AWS", "infoblox.172_28_83_29", "AUTO_CREATE", "AUTO_CREATE", "ap-northeast-1", "aws_access_key", "aws_secret_key", true, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auto_create_dns_record_type", "HOST_RECORD"),
				),
			},
			// Update and Read
			{
				Config: testAccVdiscoverytaskAutoCreateDnsRecordType(name, "A_PTR_RECORD", true, true, true, true, "$${vm_name}.testdomain.com", "AWS", "infoblox.172_28_83_29", "AUTO_CREATE", "AUTO_CREATE", "ap-northeast-1", "aws_access_key", "aws_secret_key", true, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "auto_create_dns_record_type", "A_PTR_RECORD"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVdiscoverytaskResource_CdiscoveryFileToken(t *testing.T) {
	var resourceName = "nios_discovery_vdiscovery_task.test_aws_cdiscovery_file"
	var v discovery.Vdiscoverytask

	name := acctest.RandomNameWithPrefix("example-vdiscoverytask-")

	// Get test data path
	testDataPath := getTestDataPath()
	cdiscoveryFile1 := filepath.Join(testDataPath, "cdiscoveryfile1_aws.csv")
	cdiscoveryFile2 := filepath.Join(testDataPath, "cdiscoveryfile2_aws.csv")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVdiscoverytaskCdiscoveryFile(name, "AWS", "infoblox.172_28_83_29", cdiscoveryFile1, "UPLOAD", "aws_access_key", "aws_secret_key", "arn:aws:iam::123456789012:role/InfobloxDiscoveryRole", true, true, false, true, true, true, "AUTO_CREATE", "AUTO_CREATE", true, "us-east-1", "AWS CDiscovery file test"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cdiscovery_file", cdiscoveryFile1),
				),
			},
			// Update and Read
			{
				Config: testAccVdiscoverytaskCdiscoveryFile(name, "AWS", "infoblox.172_28_83_29", cdiscoveryFile2, "UPLOAD", "aws_access_key", "aws_secret_key", "arn:aws:iam::123456789012:role/UpdatedInfobloxRole", false, false, true, true, true, true, "AUTO_CREATE", "AUTO_CREATE", false, "us-west-1", "Updated AWS CDiscovery file test"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "cdiscovery_file", cdiscoveryFile2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVdiscoverytaskResource_Comment(t *testing.T) {
	var resourceName = "nios_discovery_vdiscovery_task.test_comment"
	var v discovery.Vdiscoverytask

	name := acctest.RandomNameWithPrefix("example-vdiscoverytask-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVdiscoverytaskComment(name, "This is a test comment", true, true, true, "infoblox.172_28_83_29", "AWS", "AUTO_CREATE", "AUTO_CREATE", true, false, "ap-northeast-1", "aws_access_key", "aws_secret_key"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is a test comment"),
				),
			},
			// Update and Read
			{
				Config: testAccVdiscoverytaskComment(name, "This is an updated comment", true, true, true, "infoblox.172_28_83_29", "AWS", "AUTO_CREATE", "AUTO_CREATE", true, false, "ap-northeast-1", "aws_access_key", "aws_secret_key"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "This is an updated comment"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVdiscoverytaskResource_CredentialsType(t *testing.T) {
	var resourceName = "nios_discovery_vdiscovery_task.test_credentials_type"
	var v discovery.Vdiscoverytask

	name := acctest.RandomNameWithPrefix("example-vdiscoverytask-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVdiscoverytaskCredentialsTypeIndirect(name, "INDIRECT", "arn:aws:iam::123456789012:role/InfobloxDiscoveryRole", "DISCOVER", "infoblox.172_28_83_29", true, true, true, "AWS", "AUTO_CREATE", "AUTO_CREATE", true, true, true, "us-east-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "credentials_type", "INDIRECT"),
				),
			},
			// Update and Read
			{
				Config: testAccVdiscoverytaskCredentialsTypeDirect(name, "DIRECT", "aws_access_key", "aws_secret_key", "infoblox.172_28_83_29", true, true, true, "AWS", "AUTO_CREATE", "AUTO_CREATE", true, false, "us-east-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "credentials_type", "DIRECT"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVdiscoverytaskResource_DnsViewPrivateIp(t *testing.T) {
	var resourceName = "nios_discovery_vdiscovery_task.test_dns_view_private_ip"
	var v discovery.Vdiscoverytask

	name := acctest.RandomNameWithPrefix("example-vdiscoverytask-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVdiscoverytaskDnsViewPrivateIp(name, "default", true, "$${vm_name}.testdomain.com", "aws_access_key", "aws_secret_key", "infoblox.172_28_83_29", true, true, true, "AWS", "DIRECT", "default", "AUTO_CREATE", true, false, "us-east-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns_view_private_ip", "default"),
				),
			},
			// Update and Read
			{
				Config: testAccVdiscoverytaskDnsViewPrivateIp(name, "custom_dns_view", true, "$${vm_name}.testdomain.com", "aws_access_key", "aws_secret_key", "infoblox.172_28_83_29", true, true, true, "AWS", "DIRECT", "default", "AUTO_CREATE", true, false, "us-east-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns_view_private_ip", "custom_dns_view"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVdiscoverytaskResource_DnsViewPublicIp(t *testing.T) {
	var resourceName = "nios_discovery_vdiscovery_task.test_dns_view_public_ip"
	var v discovery.Vdiscoverytask

	name := acctest.RandomNameWithPrefix("example-vdiscoverytask-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVdiscoverytaskDnsViewPublicIp(name, "default", true, "$${vm_name}.testdomain.com", "aws_access_key", "aws_secret_key", "infoblox.172_28_83_29", true, true, true, "AWS", "AUTO_CREATE", "DIRECT", "default", true, false, "us-east-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns_view_public_ip", "default"),
				),
			},
			// Update and Read
			{
				Config: testAccVdiscoverytaskDnsViewPublicIp(name, "custom_dns_view", true, "$${vm_name}.testdomain.com", "aws_access_key", "aws_secret_key", "infoblox.172_28_83_29", true, true, true, "AWS", "AUTO_CREATE", "DIRECT", "default", true, false, "us-east-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dns_view_public_ip", "custom_dns_view"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVdiscoverytaskResource_DomainName(t *testing.T) {
	var resourceName = "nios_discovery_vdiscovery_task.test_domain_name"
	var v discovery.Vdiscoverytask

	name := acctest.RandomNameWithPrefix("example-vdiscoverytask-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVdiscoverytaskDomainName(name, "default", "openstack.example.com", "infoblox.172_28_83_29", "KEYSTONE_V3", "openstack_user", "openstack_password", true, true, true, true, "OPENSTACK", "AUTO_CREATE", "AUTO_CREATE", true, false, 443, "HTTPS"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain_name", "default"),
				),
			},
			// Update and Read
			{
				Config: testAccVdiscoverytaskDomainName(name, "custom", "openstack.example.com", "infoblox.172_28_83_29", "KEYSTONE_V3", "openstack_user", "openstack_password", true, true, true, true, "OPENSTACK", "AUTO_CREATE", "AUTO_CREATE", true, false, 443, "HTTPS"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "domain_name", "custom"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVdiscoverytaskResource_DriverType(t *testing.T) {
	var resourceName = "nios_discovery_vdiscovery_task.test_driver_type"
	var v discovery.Vdiscoverytask

	name := acctest.RandomNameWithPrefix("example-vdiscoverytask-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVdiscoverytaskDriverTypeOpenstack(name, "OPENSTACK", "openstack.example.com", "infoblox.172_28_83_29", "KEYSTONE_V2", "openstack_user", "openstack_password", true, true, true, true, "AUTO_CREATE", "AUTO_CREATE", true, false, 80, "HTTP"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "driver_type", "OPENSTACK"),
				),
			},
			// Update and Read
			{
				Config: testAccVdiscoverytaskDriverTypeVmware(name, "VMWARE", "vcenter.example.com", "infoblox.172_28_83_29", "vc_admin", "vmware_password", false, true, true, true, "AUTO_CREATE", "AUTO_CREATE", true, false, 80, "HTTP"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "driver_type", "VMWARE"),
					resource.TestCheckResourceAttr(resourceName, "fqdn_or_ip", "vcenter.example.com"),
					resource.TestCheckResourceAttr(resourceName, "port", "80"),
					resource.TestCheckResourceAttr(resourceName, "protocol", "HTTP"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVdiscoverytaskResource_EnableFilter(t *testing.T) {
	var resourceName = "nios_discovery_vdiscovery_task.test_enable_filter"
	var v discovery.Vdiscoverytask

	name := acctest.RandomNameWithPrefix("example-vdiscoverytask-")
	networklist := []string{"10.0.0.0/8", "20.0.0.0/16"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVdiscoverytaskEnableFilter(name, true, networklist, "INCLUDE", "aws_access_key", "aws_secret_key", "infoblox.172_28_83_29", true, true, true, "AWS", "AUTO_CREATE", "AUTO_CREATE", true, false, "us-east-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_filter", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccVdiscoverytaskEnableFilter(name, false, networklist, "INCLUDE", "aws_access_key", "aws_secret_key", "infoblox.172_28_83_29", true, true, true, "AWS", "AUTO_CREATE", "AUTO_CREATE", true, false, "us-east-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enable_filter", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVdiscoverytaskResource_Enabled(t *testing.T) {
	var resourceName = "nios_discovery_vdiscovery_task.test_enabled"
	var v discovery.Vdiscoverytask

	name := acctest.RandomNameWithPrefix("vdiscoverytask-enabled")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVdiscoverytaskEnabled(name, false, "azure_client_id", "azure_client_secret", "tenant_id", "infoblox.172_28_83_29", "AZURE", true, true, true, "AUTO_CREATE", "AUTO_CREATE", true, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enabled", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccVdiscoverytaskEnabled(name, true, "azure_client_id", "azure_client_secret", "tenant_id", "infoblox.172_28_83_29", "AZURE", true, true, true, "AUTO_CREATE", "AUTO_CREATE", true, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVdiscoverytaskResource_FqdnOrIp(t *testing.T) {
	var resourceName = "nios_discovery_vdiscovery_task.test_fqdn_or_ip"
	var v discovery.Vdiscoverytask

	name := acctest.RandomNameWithPrefix("example-vdiscoverytask-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVdiscoverytaskFqdnOrIp(name, "vcenter.example.com", "vc_admin", "vmware_password", "infoblox.172_28_83_29", true, true, true, true, false, "VMWARE", "AUTO_CREATE", "HTTPS", "AUTO_CREATE", "us-east-1", 443),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fqdn_or_ip", "vcenter.example.com"),
				),
			},
			// Update and Read
			{
				Config: testAccVdiscoverytaskFqdnOrIp(name, "vcenter2.example.com", "vc_admin", "vmware_password", "infoblox.172_28_83_29", true, true, true, true, false, "VMWARE", "AUTO_CREATE", "HTTPS", "AUTO_CREATE", "us-east-1", 443),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fqdn_or_ip", "vcenter2.example.com"),
				),
			},
			//Update and Read
			{
				Config: testAccVdiscoverytaskFqdnOrIp(name, "15.0.0.1", "vc_admin", "vmware_password", "infoblox.172_28_83_29", true, true, true, true, false, "VMWARE", "AUTO_CREATE", "HTTPS", "AUTO_CREATE", "us-east-1", 443),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "fqdn_or_ip", "15.0.0.1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

// "Cannot change the job type(GovCloud) during update."
func TestAccVdiscoverytaskResource_GovcloudEnabled(t *testing.T) {
	var resourceName = "nios_discovery_vdiscovery_task.test_govcloud_enabled"
	var v discovery.Vdiscoverytask

	name := acctest.RandomNameWithPrefix("example-vdiscoverytask-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVdiscoverytaskGovcloudEnabled(name, true, "aws_access_key", "aws_secret_key", "infoblox.172_28_83_29", true, true, true, "AWS", "AUTO_CREATE", "AUTO_CREATE", true, false, "us-gov-east-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "govcloud_enabled", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVdiscoverytaskResource_IdentityVersion(t *testing.T) {
	var resourceName = "nios_discovery_vdiscovery_task.test_identity_version"
	var v discovery.Vdiscoverytask

	name := acctest.RandomNameWithPrefix("example-vdiscoverytask-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVdiscoverytaskIdentityVersionV2(name, "KEYSTONE_V2", "openstack.example.com", "infoblox.172_28_83_29", "openstack_user", "openstack_password", true, true, true, true, "OPENSTACK", "AUTO_CREATE", "AUTO_CREATE", true, false, 80, "HTTP"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "identity_version", "KEYSTONE_V2"),
				),
			},
			// Update and Read
			{
				Config: testAccVdiscoverytaskIdentityVersionV3(name, "KEYSTONE_V3", "default", "openstack.example.com", "infoblox.172_28_83_29", "openstack_user", "openstack_password", true, true, true, true, "OPENSTACK", "AUTO_CREATE", "AUTO_CREATE", true, false, 80, "HTTP"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "identity_version", "KEYSTONE_V3"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVdiscoverytaskResource_Member(t *testing.T) {
	var resourceName = "nios_discovery_vdiscovery_task.test_member"
	var v discovery.Vdiscoverytask

	name := acctest.RandomNameWithPrefix("example-vdiscoverytask-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVdiscoverytaskMember(name, "infoblox.172_28_83_29", true, true, true, "AWS", "AUTO_CREATE", "AUTO_CREATE", true, false, "us-east-1", "aws_access_key", "aws_secret_key"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "member", "infoblox.172_28_83_29"),
				),
			},
			// Update and Read
			{
				Config: testAccVdiscoverytaskMember(name, "infoblox.172_28_82_115", true, true, true, "AWS", "AUTO_CREATE", "AUTO_CREATE", true, false, "us-east-1", "aws_access_key", "aws_secret_key"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "member", "infoblox.172_28_82_115"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVdiscoverytaskResource_MergeData(t *testing.T) {
	var resourceName = "nios_discovery_vdiscovery_task.test_merge_data"
	var v discovery.Vdiscoverytask

	name := acctest.RandomNameWithPrefix("example-vdiscoverytask-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVdiscoverytaskMergeData(name, true, "aws_access_key", "aws_secret_key", "infoblox.172_28_83_29", true, true, true, "AWS", "AUTO_CREATE", "AUTO_CREATE", "us-east-1", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "merge_data", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccVdiscoverytaskMergeData(name, false, "aws_access_key", "aws_secret_key", "infoblox.172_28_83_29", true, true, true, "AWS", "AUTO_CREATE", "AUTO_CREATE", "us-east-1", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "merge_data", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVdiscoverytaskResource_MultipleAccountsSyncPolicy(t *testing.T) {
	var resourceName = "nios_discovery_vdiscovery_task.test_multiple_accounts_sync_policy"
	var v discovery.Vdiscoverytask

	name := acctest.RandomNameWithPrefix("example-vdiscoverytask-")
	testDataPath := getTestDataPath()
	cdiscoveryFile := filepath.Join(testDataPath, "cdiscoveryfile1_aws.csv")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVdiscoverytaskMultipleAccountsSyncPolicyDiscover(name, "DISCOVER", true, "arn:aws:iam::123456789012:role/InfobloxDiscoveryRole", "aws_access_key", "aws_secret_key", "infoblox.172_28_83_140", true, true, true, "AWS", "AUTO_CREATE", "AUTO_CREATE", true, true, "us-east-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "multiple_accounts_sync_policy", "DISCOVER"),
				),
			},
			// Update and Read
			{
				Config: testAccVdiscoverytaskMultipleAccountsSyncPolicyUpload(name, "UPLOAD", true, cdiscoveryFile, "arn:aws:iam::123456789012:role/InfobloxDiscoveryRole", "aws_access_key", "aws_secret_key", "infoblox.172_28_83_140", true, true, true, "AWS", "AUTO_CREATE", "AUTO_CREATE", true, true, "us-east-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "multiple_accounts_sync_policy", "UPLOAD"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVdiscoverytaskResource_Name(t *testing.T) {
	var resourceName = "nios_discovery_vdiscovery_task.test_name"
	var v discovery.Vdiscoverytask

	name1 := acctest.RandomNameWithPrefix("vdiscoverytask-")
	name2 := acctest.RandomNameWithPrefix("updated-vdiscoverytask-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVdiscoverytaskName(name1, "infoblox.172_28_83_29", true, true, true, "AWS", "AUTO_CREATE", "AUTO_CREATE", true, false, "us-east-1", "aws_access_key", "aws_secret_key"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name1),
				),
			},
			// Update and Read
			{
				Config: testAccVdiscoverytaskName(name2, "infoblox.172_28_83_29", true, true, true, "AWS", "AUTO_CREATE", "AUTO_CREATE", true, false, "us-east-1", "aws_access_key", "aws_secret_key"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", name2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVdiscoverytaskResource_NetworkFilter(t *testing.T) {
	var resourceName = "nios_discovery_vdiscovery_task.test_network_filter"
	var v discovery.Vdiscoverytask

	name := acctest.RandomNameWithPrefix("example-vdiscoverytask-")
	networkList := []string{"10.0.0.0/8", "25.0.0.0/16"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVdiscoverytaskNetworkFilter(name, true, "INCLUDE", networkList, "aws_access_key", "aws_secret_key", "infoblox.172_28_83_29", true, true, true, "AWS", "AUTO_CREATE", "AUTO_CREATE", true, false, "us-east-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_filter", "INCLUDE"),
				),
			},
			// Update and Read
			{
				Config: testAccVdiscoverytaskNetworkFilter(name, true, "EXCLUDE", networkList, "aws_access_key", "aws_secret_key", "infoblox.172_28_83_29", true, true, true, "AWS", "AUTO_CREATE", "AUTO_CREATE", true, false, "us-east-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_filter", "EXCLUDE"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVdiscoverytaskResource_NetworkList(t *testing.T) {
	var resourceName = "nios_discovery_vdiscovery_task.test_network_list"
	var v discovery.Vdiscoverytask

	name := acctest.RandomNameWithPrefix("example-vdiscoverytask-")
	networkList1 := []string{"10.0.0.0/8", "192.168.0.0/16"}
	networkList2 := []string{"172.16.0.0/12", "203.0.113.0/24", "198.51.100.0/24"}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVdiscoverytaskNetworkList(name, networkList1, "aws_access_key", "aws_secret_key", "infoblox.172_28_83_29", true, true, true, "AWS", "AUTO_CREATE", "AUTO_CREATE", true, false, "us-east-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_list.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "network_list.0", "10.0.0.0/8"),
					resource.TestCheckResourceAttr(resourceName, "network_list.1", "192.168.0.0/16"),
				),
			},
			// Update and Read
			{
				Config: testAccVdiscoverytaskNetworkList(name, networkList2, "aws_access_key", "aws_secret_key", "infoblox.172_28_83_29", true, true, true, "AWS", "AUTO_CREATE", "AUTO_CREATE", true, false, "us-east-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "network_list.#", "3"),
					resource.TestCheckResourceAttr(resourceName, "network_list.0", "172.16.0.0/12"),
					resource.TestCheckResourceAttr(resourceName, "network_list.1", "203.0.113.0/24"),
					resource.TestCheckResourceAttr(resourceName, "network_list.2", "198.51.100.0/24"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVdiscoverytaskResource_Password(t *testing.T) {
	var resourceName = "nios_discovery_vdiscovery_task.test_password"
	var v discovery.Vdiscoverytask

	name := acctest.RandomNameWithPrefix("example-vdiscoverytask-")
	password1 := "aws_secret_key1"
	password2 := "aws_secret_key2"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVdiscoverytaskPassword(name, "aws_access_key", password1, "infoblox.172_28_83_29", true, true, true, "AWS", "AUTO_CREATE", "AUTO_CREATE", true, false, "us-east-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "password", password1),
				),
			},
			// Update and Read
			{
				Config: testAccVdiscoverytaskPassword(name, "aws_access_key", password2, "infoblox.172_28_83_29", true, true, true, "AWS", "AUTO_CREATE", "AUTO_CREATE", true, false, "us-east-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "password", password2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVdiscoverytaskResource_Port(t *testing.T) {
	var resourceName = "nios_discovery_vdiscovery_task.test_port"
	var v discovery.Vdiscoverytask

	name := acctest.RandomNameWithPrefix("example-vdiscoverytask-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVdiscoverytaskPort(name, 443, "vc_admin", "vmware_password", "infoblox.172_28_83_29", "vcenter.example.com", true, true, true, true, false, "VMWARE", "AUTO_CREATE", "HTTPS", "AUTO_CREATE", "us-east-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port", "443"),
				),
			},
			// Update and Read
			{
				Config: testAccVdiscoverytaskPort(name, 8080, "vc_admin", "vmware_password", "infoblox.172_28_83_29", "vcenter.example.com", true, true, true, true, false, "VMWARE", "AUTO_CREATE", "HTTPS", "AUTO_CREATE", "us-east-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "port", "8080"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVdiscoverytaskResource_PrivateNetworkView(t *testing.T) {
	var resourceName = "nios_discovery_vdiscovery_task.test_private_network_view"
	var v discovery.Vdiscoverytask

	name := acctest.RandomNameWithPrefix("example-vdiscoverytask-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVdiscoverytaskPrivateNetworkView(name, "default", "DIRECT", "aws_access_key", "aws_secret_key", "infoblox.172_28_83_29", true, true, true, "AWS", "AUTO_CREATE", true, false, "us-east-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "private_network_view", "default"),
				),
			},
			// Update and Read
			{
				Config: testAccVdiscoverytaskPrivateNetworkView(name, "custom_private_view", "DIRECT", "aws_access_key", "aws_secret_key", "infoblox.172_28_83_29", true, true, true, "AWS", "AUTO_CREATE", true, false, "us-east-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "private_network_view", "custom_private_view"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccVdiscoverytaskResource_PrivateNetworkViewMappingPolicy(t *testing.T) {
	var resourceName = "nios_discovery_vdiscovery_task.test_private_network_view_mapping_policy"
	var v discovery.Vdiscoverytask

	name := acctest.RandomNameWithPrefix("example-vdiscoverytask-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVdiscoverytaskPrivateNetworkViewMappingPolicyAutoCreate(name, "AUTO_CREATE", "aws_access_key", "aws_secret_key", "infoblox.172_28_83_29", true, true, true, "AWS", "AUTO_CREATE", true, false, "us-east-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "private_network_view_mapping_policy", "AUTO_CREATE"),
				),
			},
			// Update and Read
			{
				Config: testAccVdiscoverytaskPrivateNetworkViewMappingPolicyDirect(name, "DIRECT", "default", "aws_access_key", "aws_secret_key", "infoblox.172_28_83_29", true, true, true, "AWS", "AUTO_CREATE", true, false, "us-east-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "private_network_view_mapping_policy", "DIRECT"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVdiscoverytaskResource_Protocol(t *testing.T) {
	var resourceName = "nios_discovery_vdiscovery_task.test_protocol"
	var v discovery.Vdiscoverytask

	name := acctest.RandomNameWithPrefix("example-vdiscoverytask-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVdiscoverytaskProtocol(name, "HTTPS", "vc_admin", "vmware_password", "infoblox.172_28_83_29", "vcenter.example.com", true, true, true, true, false, "VMWARE", "AUTO_CREATE", "AUTO_CREATE", "us-east-1", 443),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "protocol", "HTTPS"),
				),
			},
			// Update and Read
			{
				Config: testAccVdiscoverytaskProtocol(name, "HTTP", "vc_admin", "vmware_password", "infoblox.172_28_83_29", "vcenter.example.com", true, true, true, true, false, "VMWARE", "AUTO_CREATE", "AUTO_CREATE", "us-east-1", 443),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "protocol", "HTTP"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVdiscoverytaskResource_PublicNetworkView(t *testing.T) {
	var resourceName = "nios_discovery_vdiscovery_task.test_public_network_view"
	var v discovery.Vdiscoverytask

	name := acctest.RandomNameWithPrefix("example-vdiscoverytask-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVdiscoverytaskPublicNetworkView(name, "default", "DIRECT", "aws_access_key", "aws_secret_key", "infoblox.172_28_83_29", true, true, true, "AWS", "AUTO_CREATE", true, false, "us-east-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "public_network_view", "default"),
				),
			},
			// Update and Read
			{
				Config: testAccVdiscoverytaskPublicNetworkView(name, "custom_public_view", "DIRECT", "aws_access_key", "aws_secret_key", "infoblox.172_28_83_29", true, true, true, "AWS", "AUTO_CREATE", true, false, "us-east-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "public_network_view", "custom_public_view"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVdiscoverytaskResource_PublicNetworkViewMappingPolicy(t *testing.T) {
	var resourceName = "nios_discovery_vdiscovery_task.test_public_network_view_mapping_policy"
	var v discovery.Vdiscoverytask

	name := acctest.RandomNameWithPrefix("example-vdiscoverytask-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVdiscoverytaskPublicNetworkViewMappingPolicyAutoCreate(name, "AUTO_CREATE", "aws_access_key", "aws_secret_key", "infoblox.172_28_83_29", true, true, true, "AWS", "AUTO_CREATE", true, false, "us-east-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "public_network_view_mapping_policy", "AUTO_CREATE"),
				),
			},
			// Update and Read
			{
				Config: testAccVdiscoverytaskPublicNetworkViewMappingPolicyDirect(name, "DIRECT", "default", "aws_access_key", "aws_secret_key", "infoblox.172_28_83_29", true, true, true, "AWS", "AUTO_CREATE", true, false, "us-east-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "public_network_view_mapping_policy", "DIRECT"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVdiscoverytaskResource_RoleArn(t *testing.T) {
	var resourceName = "nios_discovery_vdiscovery_task.test_role_arn"
	var v discovery.Vdiscoverytask

	name := acctest.RandomNameWithPrefix("example-vdiscoverytask-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVdiscoverytaskRoleArn(name, "arn:aws:iam::123456789012:role/InfobloxDiscoveryRole", "DISCOVER", true, "aws_access_key", "aws_secret_key", "infoblox.172_28_83_140", true, true, true, "AWS", "AUTO_CREATE", "AUTO_CREATE", true, true, "us-east-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "role_arn", "arn:aws:iam::123456789012:role/InfobloxDiscoveryRole"),
				),
			},
			// Update and Read
			{
				Config: testAccVdiscoverytaskRoleArn(name, "arn:aws:iam::123456789012:role/UpdatedInfobloxRole", "DISCOVER", true, "aws_access_key", "aws_secret_key", "infoblox.172_28_83_140", true, true, true, "AWS", "AUTO_CREATE", "AUTO_CREATE", true, true, "us-east-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "role_arn", "arn:aws:iam::123456789012:role/UpdatedInfobloxRole"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVdiscoverytaskResource_ScheduledRun(t *testing.T) {
	var resourceName = "nios_discovery_vdiscovery_task.test_scheduled_run_block_azure"
	var v discovery.Vdiscoverytask

	name := acctest.RandomNameWithPrefix("vdiscoverytask-azure-schedule")

	scheduledRun1 := map[string]any{
		"time_zone":         "UTC",
		"weekdays":          []string{"MONDAY", "THURSDAY"},
		"frequency":         "WEEKLY",
		"every":             1,
		"minutes_past_hour": 30,
		"hour_of_day":       9,
		"repeat":            "RECUR",
		"disable":           false,
	}

	scheduledRun2 := map[string]any{
		"time_zone":         "Asia/Kolkata",
		"day_of_month":      21,
		"month":             12,
		"year":              2025,
		"minutes_past_hour": 45,
		"hour_of_day":       22,
		"repeat":            "ONCE",
		"disable":           true,
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVdiscoverytaskScheduledRun(name, scheduledRun1, "azure_client_id", "azure_client_secret", "tenant_id", "infoblox.172_28_83_29", "AZURE", true, true, true, "AUTO_CREATE", "AUTO_CREATE", true, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "scheduled_run.frequency", "WEEKLY"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_run.time_zone", "UTC"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_run.hour_of_day", "9"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_run.repeat", "RECUR"),
				),
			},
			// Update and Read
			{
				Config: testAccVdiscoverytaskScheduledRun(name, scheduledRun2, "azure_client_id", "azure_client_secret", "tenant_id", "infoblox.172_28_83_29", "AZURE", true, true, true, "AUTO_CREATE", "AUTO_CREATE", true, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "scheduled_run.time_zone", "Asia/Kolkata"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_run.hour_of_day", "22"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_run.day_of_month", "21"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_run.repeat", "ONCE"),
					resource.TestCheckResourceAttr(resourceName, "scheduled_run.disable", "true"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVdiscoverytaskResource_SelectedRegions(t *testing.T) {
	var resourceName = "nios_discovery_vdiscovery_task.test_selected_regions"
	var v discovery.Vdiscoverytask

	name := acctest.RandomNameWithPrefix("example-vdiscoverytask-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVdiscoverytaskSelectedRegions(name, "us-east-1", "aws_access_key", "aws_secret_key", "infoblox.172_28_83_29", true, true, true, "AWS", "AUTO_CREATE", "AUTO_CREATE", true, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "selected_regions", "us-east-1"),
				),
			},
			// Update and Read
			{
				Config: testAccVdiscoverytaskSelectedRegions(name, "us-west-1", "aws_access_key", "aws_secret_key", "infoblox.172_28_83_29", true, true, true, "AWS", "AUTO_CREATE", "AUTO_CREATE", true, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "selected_regions", "us-west-1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVdiscoverytaskResource_ServiceAccountFile(t *testing.T) {
	var resourceName = "nios_discovery_vdiscovery_task.test_gcp_service_account_file"
	var v discovery.Vdiscoverytask

	name := acctest.RandomNameWithPrefix("example-vdiscoverytask-")

	// Get test data path
	testDataPath := getTestDataPath()
	serviceAccountFile1 := filepath.Join(testDataPath, "service_account_file1.json")
	serviceAccountFile2 := filepath.Join(testDataPath, "service_account_file2.json")
	cdiscoveryFile := filepath.Join(testDataPath, "cdiscoveryfile_gcp.csv")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVdiscoverytaskServiceAccountFile(name, "GCP", "infoblox.172_28_83_140", serviceAccountFile1, cdiscoveryFile, "DISCOVER", true, true, true, true, false, "AUTO_CREATE", "AUTO_CREATE", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "service_account_file", serviceAccountFile1),
					resource.TestCheckResourceAttr(resourceName, "cdiscovery_file", cdiscoveryFile),
					resource.TestCheckResourceAttr(resourceName, "multiple_accounts_sync_policy", "DISCOVER"),
				),
			},
			// Update and Read
			{
				Config: testAccVdiscoverytaskServiceAccountFile(name, "GCP", "infoblox.172_28_83_140", serviceAccountFile2, cdiscoveryFile, "DISCOVER", false, false, true, true, true, "AUTO_CREATE", "AUTO_CREATE", false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "service_account_file", serviceAccountFile2),
					resource.TestCheckResourceAttr(resourceName, "cdiscovery_file", cdiscoveryFile),
					resource.TestCheckResourceAttr(resourceName, "multiple_accounts_sync_policy", "DISCOVER"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccVdiscoverytaskResource_SyncChildAccounts(t *testing.T) {
	var resourceName = "nios_discovery_vdiscovery_task.test_sync_child_accounts"
	var v discovery.Vdiscoverytask

	name := acctest.RandomNameWithPrefix("example-vdiscoverytask-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVdiscoverytaskSyncChildAccounts(name, true, "arn:aws:iam::123456789012:role/InfobloxDiscoveryRole", "DISCOVER", "aws_access_key", "aws_secret_key", "infoblox.172_28_83_140", true, true, true, "AWS", "AUTO_CREATE", "AUTO_CREATE", true, true, "us-east-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "sync_child_accounts", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccVdiscoverytaskSyncChildAccounts(name, false, "arn:aws:iam::123456789012:role/UpdatedInfobloxRole", "DISCOVER", "aws_access_key", "aws_secret_key", "infoblox.172_28_83_140", true, true, true, "AWS", "AUTO_CREATE", "AUTO_CREATE", true, true, "us-east-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "sync_child_accounts", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVdiscoverytaskResource_UpdateDnsViewPrivateIp(t *testing.T) {
	var resourceName = "nios_discovery_vdiscovery_task.test_update_dns_view_private_ip"
	var v discovery.Vdiscoverytask

	name := acctest.RandomNameWithPrefix("example-vdiscoverytask-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccVdiscoverytaskUpdateDnsViewPrivateIpFalse(name, false, "aws_access_key", "aws_secret_key", "infoblox.172_28_83_29", true, true, true, "AWS", "AUTO_CREATE", "AUTO_CREATE", true, false, "us-east-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_dns_view_private_ip", "false"),
				),
			},
			{
				Config: testAccVdiscoverytaskUpdateDnsViewPrivateIpTrue(name, true, "custom_dns_view", "aws_access_key", "aws_secret_key", "infoblox.172_28_83_29", true, true, true, "AWS", "DIRECT", "default", "AUTO_CREATE", true, false, "us-east-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_dns_view_private_ip", "true"),
					resource.TestCheckResourceAttr(resourceName, "dns_view_private_ip", "custom_dns_view"),
				),
			},
		},
	})
}

func TestAccVdiscoverytaskResource_UpdateDnsViewPublicIp(t *testing.T) {
	var resourceName = "nios_discovery_vdiscovery_task.test_update_dns_view_public_ip"
	var v discovery.Vdiscoverytask

	name := acctest.RandomNameWithPrefix("example-vdiscoverytask-public")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			//Create and Read
			{
				Config: testAccVdiscoverytaskUpdateDnsViewPublicIpFalse(name, false, "aws_access_key", "aws_secret_key", "infoblox.172_28_83_29", true, true, true, "AWS", "AUTO_CREATE", "AUTO_CREATE", true, false, "us-east-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_dns_view_public_ip", "false"),
				),
			},
			// Update and Read
			{
				Config: testAccVdiscoverytaskUpdateDnsViewPublicIpTrue(name, true, "custom_dns_view", "aws_access_key", "aws_secret_key", "infoblox.172_28_83_29", true, true, true, "AWS", "AUTO_CREATE", "DIRECT", "default", true, false, "us-east-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_dns_view_public_ip", "true"),
					resource.TestCheckResourceAttr(resourceName, "dns_view_public_ip", "custom_dns_view"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccVdiscoverytaskResource_UpdateMetadata(t *testing.T) {
	var resourceName = "nios_discovery_vdiscovery_task.test_update_metadata"
	var v discovery.Vdiscoverytask

	name := acctest.RandomNameWithPrefix("example-vdiscoverytask-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVdiscoverytaskUpdateMetadata(name, true, "aws_access_key", "aws_secret_key", "infoblox.172_28_83_29", true, true, true, "AWS", "AUTO_CREATE", "AUTO_CREATE", "us-east-1", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_metadata", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccVdiscoverytaskUpdateMetadata(name, false, "aws_access_key", "aws_secret_key", "infoblox.172_28_83_29", true, true, true, "AWS", "AUTO_CREATE", "AUTO_CREATE", "us-east-1", true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "update_metadata", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccVdiscoverytaskResource_UseIdentity(t *testing.T) {
	var resourceName = "nios_discovery_vdiscovery_task.test_use_identity"
	var v discovery.Vdiscoverytask

	name := acctest.RandomNameWithPrefix("example-vdiscoverytask-")

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVdiscoverytaskUseIdentity(name, true, "openstack.example.com", 80, "HTTP", "infoblox.172_28_83_29", "KEYSTONE_V2", "openstack_user", "openstack_password", true, true, true, "OPENSTACK", "AUTO_CREATE", "AUTO_CREATE", true, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_identity", "true"),
				),
			},
			// Update and Read
			{
				Config: testAccVdiscoverytaskUseIdentity(name, false, "openstack.example.com", 80, "HTTP", "infoblox.172_28_83_29", "KEYSTONE_V2", "openstack_user", "openstack_password", true, true, true, "OPENSTACK", "AUTO_CREATE", "AUTO_CREATE", true, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "use_identity", "false"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
func TestAccVdiscoverytaskResource_Username(t *testing.T) {
	var resourceName = "nios_discovery_vdiscovery_task.test_username"
	var v discovery.Vdiscoverytask

	name := acctest.RandomNameWithPrefix("example-vdiscoverytask-")
	username1 := "User1"
	username2 := "User2"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccVdiscoverytaskUsername(name, username1, "aws_secret_key", "infoblox.172_28_83_29", true, true, true, "AWS", "AUTO_CREATE", "AUTO_CREATE", true, false, "us-east-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "username", username1),
				),
			},
			// Update and Read
			{
				Config: testAccVdiscoverytaskUsername(name, username2, "aws_secret_key", "infoblox.172_28_83_29", true, true, true, "AWS", "AUTO_CREATE", "AUTO_CREATE", true, false, "us-east-1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVdiscoverytaskExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "username", username2),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckVdiscoverytaskExists(ctx context.Context, resourceName string, v *discovery.Vdiscoverytask) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.NIOSClient.DiscoveryAPI.
			VdiscoverytaskAPI.
			Read(ctx, utils.ExtractResourceRef(rs.Primary.Attributes["ref"])).
			ReturnFieldsPlus(readableAttributesForVdiscoverytask).
			ReturnAsObject(1).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.GetVdiscoverytaskResponseObjectAsResult.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetVdiscoverytaskResponseObjectAsResult.GetResult()
		return nil
	}
}

func testAccCheckVdiscoverytaskDestroy(ctx context.Context, v *discovery.Vdiscoverytask) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.NIOSClient.DiscoveryAPI.
			VdiscoverytaskAPI.
			Read(ctx, utils.ExtractResourceRef(*v.Ref)).
			ReturnAsObject(1).
			ReturnFieldsPlus(readableAttributesForVdiscoverytask).
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

func testAccCheckVdiscoverytaskDisappears(ctx context.Context, v *discovery.Vdiscoverytask) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.NIOSClient.DiscoveryAPI.
			VdiscoverytaskAPI.
			Delete(ctx, utils.ExtractResourceRef(*v.Ref)).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccVdiscoverytaskBasicConfig(name string, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm bool, driverType, member, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy string, mergeData, updateMetadata bool, selectedRegions, username, password string) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test" {
    name                                = %q
    driver_type                         = %q
    member                              = %q
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    merge_data                          = %t
    update_metadata                     = %t
    private_network_view_mapping_policy = %q
    public_network_view_mapping_policy  = %q
    selected_regions                    = %q
    username                            = %q
    password                            = %q
}
`, name, driverType, member, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, mergeData, updateMetadata, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, selectedRegions, username, password)
}

func testAccVdiscoverytaskAllowUnsecuredConnection(name string, allowUnsecuredConnection bool, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, mergeData, updateMetadata bool, driverType, fqdnOrIp, member, password, privateNetworkViewMappingPolicy, protocol, publicNetworkViewMappingPolicy, selectedRegions, username string, port int) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_allow_unsecured_connection" {
    name                                = %q
    driver_type                         = %q
    member                              = %q
    allow_unsecured_connection          = %t
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    merge_data                          = %t
    update_metadata                     = %t
    fqdn_or_ip                          = %q
    protocol                            = %q
    port                                = %d
    private_network_view_mapping_policy = %q
    public_network_view_mapping_policy  = %q
    selected_regions                    = %q
    username                            = %q
    password                            = %q
}
`, name, driverType, member, allowUnsecuredConnection, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, mergeData, updateMetadata, fqdnOrIp, protocol, port, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, selectedRegions, username, password)
}

func testAccVdiscoverytaskAutoConsolidateCloudEa(name string, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, mergeData, updateMetadata bool, driverType, member, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, selectedRegions, username, password string) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_auto_consolidate_cloud_ea" {
    name                                = %q
    driver_type                         = %q
    member                              = %q
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    merge_data                          = %t
    update_metadata                     = %t
    private_network_view_mapping_policy = %q
    public_network_view_mapping_policy  = %q
    selected_regions                    = %q
    username                            = %q
    password                            = %q
}
`, name, driverType, member, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, mergeData, updateMetadata, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, selectedRegions, username, password)
}

func testAccVdiscoverytaskAutoConsolidateManagedTenant(name string, autoConsolidateManagedTenant, autoConsolidateCloudEa, autoConsolidateManagedVm, mergeData, updateMetadata bool, driverType, member, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, selectedRegions, username, password string) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_auto_consolidate_managed_tenant" {
    name                                = %q
    driver_type                         = %q
    member                              = %q
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_vm         = %t
    merge_data                          = %t
    update_metadata                     = %t
    private_network_view_mapping_policy = %q
    public_network_view_mapping_policy  = %q
    selected_regions                    = %q
    username                            = %q
    password                            = %q
}
`, name, driverType, member, autoConsolidateManagedTenant, autoConsolidateCloudEa, autoConsolidateManagedVm, mergeData, updateMetadata, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, selectedRegions, username, password)
}

func testAccVdiscoverytaskAutoConsolidateManagedVm(name string, autoConsolidateManagedVm, autoConsolidateCloudEa, autoConsolidateManagedTenant, mergeData, updateMetadata bool, driverType, member, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, selectedRegions, username, password string) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_auto_consolidate_managed_vm" {
    name                                = %q
    driver_type                         = %q
    member                              = %q
    auto_consolidate_managed_vm         = %t
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    merge_data                          = %t
    update_metadata                     = %t
    private_network_view_mapping_policy = %q
    public_network_view_mapping_policy  = %q
    selected_regions                    = %q
    username                            = %q
    password                            = %q
}
`, name, driverType, member, autoConsolidateManagedVm, autoConsolidateCloudEa, autoConsolidateManagedTenant, mergeData, updateMetadata, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, selectedRegions, username, password)
}

func testAccVdiscoverytaskAutoCreateDnsHostnameTemplate(name string, autoCreateDnsHostnameTemplate string, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, autoCreateDnsRecord bool, autoCreateDnsRecordType, driverType, member, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, selectedRegions, username, password string, mergeData, updateMetadata bool) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_auto_create_dns_hostname_template" {
    name                                = %q
    driver_type                         = %q
    member                              = %q
    auto_create_dns_hostname_template   = %q
    auto_create_dns_record              = %t
    auto_create_dns_record_type         = %q
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    merge_data                          = %t
    update_metadata                     = %t
    private_network_view_mapping_policy = %q
    public_network_view_mapping_policy  = %q
    selected_regions                    = %q
    username                            = %q
    password                            = %q
}
`, name, driverType, member, autoCreateDnsHostnameTemplate, autoCreateDnsRecord, autoCreateDnsRecordType, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, mergeData, updateMetadata, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, selectedRegions, username, password)
}

func testAccVdiscoverytaskAutoCreateDnsRecord(name string, autoCreateDnsRecord bool, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm bool, autoCreateDnsHostnameTemplate, autoCreateDnsRecordType, driverType, member, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, selectedRegions, username, password string, mergeData, updateMetadata bool) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_auto_create_dns_record" {
    name                                = %q
    driver_type                         = %q
    member                              = %q
    auto_create_dns_record              = %t
    auto_create_dns_hostname_template   = %q
    auto_create_dns_record_type         = %q
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    merge_data                          = %t
    update_metadata                     = %t
    private_network_view_mapping_policy = %q
    public_network_view_mapping_policy  = %q
    selected_regions                    = %q
    username                            = %q
    password                            = %q
}
`, name, driverType, member, autoCreateDnsRecord, autoCreateDnsHostnameTemplate, autoCreateDnsRecordType, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, mergeData, updateMetadata, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, selectedRegions, username, password)
}

func testAccVdiscoverytaskAutoCreateDnsRecordType(name string, autoCreateDnsRecordType string, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, autoCreateDnsRecord bool, autoCreateDnsHostnameTemplate, driverType, member, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, selectedRegions, username, password string, mergeData, updateMetadata bool) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_auto_create_dns_record_type" {
    name                                = %q
    driver_type                         = %q
    member                              = %q
    auto_create_dns_record_type         = %q
    auto_create_dns_record              = %t
    auto_create_dns_hostname_template   = %q
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    merge_data                          = %t
    update_metadata                     = %t
    private_network_view_mapping_policy = %q
    public_network_view_mapping_policy  = %q
    selected_regions                    = %q
    username                            = %q
    password                            = %q
}
`, name, driverType, member, autoCreateDnsRecordType, autoCreateDnsRecord, autoCreateDnsHostnameTemplate, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, mergeData, updateMetadata, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, selectedRegions, username, password)
}
func testAccVdiscoverytaskCdiscoveryFile(name, driverType, member, cdiscoveryFile, multipleAccountsSyncPolicy, username, password, roleArn string, syncChildAccounts, mergeData, updateMetadata, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm bool, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy string, enabled bool, selectedRegions, comment string) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_aws_cdiscovery_file" {
    name                                = %q
    driver_type                         = %q
    member                              = %q
    cdiscovery_file                     = %q
    multiple_accounts_sync_policy       = %q
    username                            = %q
    password                            = %q
    role_arn                            = %q
    sync_child_accounts                 = %t
    merge_data                          = %t
    update_metadata                     = %t
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    private_network_view_mapping_policy = %q
    public_network_view_mapping_policy  = %q
    enabled                             = %t
    selected_regions                    = %q
    comment                             = %q
}
`, name, driverType, member, cdiscoveryFile, multipleAccountsSyncPolicy, username, password, roleArn, syncChildAccounts, mergeData, updateMetadata, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, enabled, selectedRegions, comment)
}

func testAccVdiscoverytaskComment(name, comment string, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm bool, member, driverType, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy string, mergeData, updateMetadata bool, selectedRegions, username, password string) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_comment" {
    name                                = %q
    comment                             = %q
    driver_type                         = %q
    member                              = %q
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    merge_data                          = %t
    update_metadata                     = %t
    private_network_view_mapping_policy = %q
    public_network_view_mapping_policy  = %q
    selected_regions                    = %q
    username                            = %q
    password                            = %q
}
`, name, comment, driverType, member, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, mergeData, updateMetadata, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, selectedRegions, username, password)
}

func testAccVdiscoverytaskCredentialsTypeIndirect(name, credentialsType, roleArn, multipleAccountsSyncPolicy, member string, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm bool, driverType, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy string, mergeData, syncChildAccounts, update_metadata bool, selectedRegions string) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_credentials_type" {
    name                                = %q
    credentials_type                    = %q
    role_arn                            = %q
    multiple_accounts_sync_policy       = %q
    member                              = %q
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    driver_type                         = %q
    private_network_view_mapping_policy = %q
    public_network_view_mapping_policy  = %q
    merge_data                          = %t
    sync_child_accounts                 = %t
	update_metadata                     = %t
    selected_regions                    = %q
}
`, name, credentialsType, roleArn, multipleAccountsSyncPolicy, member, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, driverType, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, mergeData, syncChildAccounts, update_metadata, selectedRegions)
}

func testAccVdiscoverytaskCredentialsTypeDirect(name, credentialsType, username, password, member string, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm bool, driverType, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy string, mergeData, updateMetadata bool, selectedRegions string) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_credentials_type" {
    name                                = %q
    credentials_type                    = %q
    username                            = %q
    password                            = %q
    member                              = %q
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    driver_type                         = %q
    private_network_view_mapping_policy = %q
    public_network_view_mapping_policy  = %q
    merge_data                          = %t
    update_metadata                     = %t
    selected_regions                    = %q
}
`, name, credentialsType, username, password, member, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, driverType, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, mergeData, updateMetadata, selectedRegions)
}
func testAccVdiscoverytaskDnsViewPrivateIp(name, dnsViewPrivateIp string, updateDnsViewPrivateIp bool, autoCreateDnsHostnameTemplate, username, password, member string, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm bool, driverType, privateNetworkViewMappingPolicy, privateNetworkView, publicNetworkViewMappingPolicy string, mergeData, updateMetadata bool, selectedRegions string) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_dns_view_private_ip" {
    name                                = %q
    dns_view_private_ip                 = %q
    update_dns_view_private_ip          = %t
    auto_create_dns_hostname_template   = %q
    username                            = %q
    password                            = %q
    member                              = %q
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    driver_type                         = %q
    private_network_view_mapping_policy = %q
    private_network_view                = %q
    public_network_view_mapping_policy  = %q
    merge_data                          = %t
    update_metadata                     = %t
    selected_regions                    = %q
}
`, name, dnsViewPrivateIp, updateDnsViewPrivateIp, autoCreateDnsHostnameTemplate, username, password, member, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, driverType, privateNetworkViewMappingPolicy, privateNetworkView, publicNetworkViewMappingPolicy, mergeData, updateMetadata, selectedRegions)
}

func testAccVdiscoverytaskDnsViewPublicIp(name, dnsViewPublicIp string, updateDnsViewPublicIp bool, autoCreateDnsHostnameTemplate, username, password, member string, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm bool, driverType, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, publicNetworkView string, mergeData, updateMetadata bool, selectedRegions string) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_dns_view_public_ip" {
    name                                = %q
    dns_view_public_ip                  = %q
    update_dns_view_public_ip           = %t
    auto_create_dns_hostname_template   = %q
    username                            = %q
    password                            = %q
    member                              = %q
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    driver_type                         = %q
    private_network_view_mapping_policy = %q
    public_network_view_mapping_policy  = %q
    public_network_view                 = %q
    merge_data                          = %t
    update_metadata                     = %t
    selected_regions                    = %q
}
`, name, dnsViewPublicIp, updateDnsViewPublicIp, autoCreateDnsHostnameTemplate, username, password, member, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, driverType, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, publicNetworkView, mergeData, updateMetadata, selectedRegions)
}

func testAccVdiscoverytaskDomainName(name, domainName, fqdnOrIp, member, identityVersion, username, password string, useIdentity, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm bool, driverType, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy string, mergeData, updateMetadata bool, port int, protocol string) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_domain_name" {
    name                                = %q
    domain_name                         = %q
    fqdn_or_ip                          = %q
    member                              = %q
    identity_version                    = %q
    username                            = %q
    password                            = %q
    use_identity                        = %t
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    driver_type                         = %q
    private_network_view_mapping_policy = %q
    public_network_view_mapping_policy  = %q
    merge_data                          = %t
    update_metadata                     = %t
    port                                = %d
    protocol                            = %q
}
`, name, domainName, fqdnOrIp, member, identityVersion, username, password, useIdentity, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, driverType, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, mergeData, updateMetadata, port, protocol)
}

func testAccVdiscoverytaskDriverTypeOpenstack(name, driverType, fqdnOrIp, member, identityVersion, username, password string, useIdentity, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm bool, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy string, mergeData, updateMetadata bool, port int, protocol string) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_driver_type" {
    name                                = %q
    driver_type                         = %q
    fqdn_or_ip                          = %q
    member                              = %q
    identity_version                    = %q
    username                            = %q
    password                            = %q
    use_identity                        = %t
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    private_network_view_mapping_policy = %q
    public_network_view_mapping_policy  = %q
    merge_data                          = %t
    update_metadata                     = %t
    port                                = %d
    protocol                            = %q
}
`, name, driverType, fqdnOrIp, member, identityVersion, username, password, useIdentity, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, mergeData, updateMetadata, port, protocol)
}

func testAccVdiscoverytaskDriverTypeVmware(name, driverType, fqdnOrIp, member, username, password string, useIdentity, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm bool, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy string, mergeData, updateMetadata bool, port int, protocol string) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_driver_type" {
    name                                = %q
    driver_type                         = %q
    fqdn_or_ip                          = %q
    member                              = %q
    username                            = %q
    password                            = %q
    use_identity                        = %t
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    private_network_view_mapping_policy = %q
    public_network_view_mapping_policy  = %q
    merge_data                          = %t
    update_metadata                     = %t
    port                                = %d
    protocol                            = %q
}
`, name, driverType, fqdnOrIp, member, username, password, useIdentity, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, mergeData, updateMetadata, port, protocol)
}

func testAccVdiscoverytaskEnableFilter(name string, enableFilter bool, networkList []string, networkFilter, username, password, member string, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm bool, driverType, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy string, mergeData, updateMetadata bool, selectedRegions string) string {
	networkListHCL := utils.ConvertStringSliceToHCL(networkList)
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_enable_filter" {
    name                                = %q
    enable_filter                       = %t
    network_list                        = %s
    network_filter                      = %q
    username                            = %q
    password                            = %q
    member                              = %q
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    driver_type                         = %q
    private_network_view_mapping_policy = %q
    public_network_view_mapping_policy  = %q
    merge_data                          = %t
    update_metadata                     = %t
    selected_regions                    = %q
}
`, name, enableFilter, networkListHCL, networkFilter, username, password, member, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, driverType, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, mergeData, updateMetadata, selectedRegions)
}

func testAccVdiscoverytaskEnabled(name string, enabled bool, username, password, fqdnOrIp, member, driverType string, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm bool, privateMappingPolicy, publicMappingPolicy string, mergeData, updateMetadata bool) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_enabled" {
    name                                = %q
    enabled                             = %t
    username                            = %q
    password                            = %q
    fqdn_or_ip                          = %q
    member                              = %q
    driver_type                         = %q
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    private_network_view_mapping_policy = %q
    public_network_view_mapping_policy  = %q
    merge_data                          = %t
    update_metadata                     = %t
}
`, name, enabled, username, password, fqdnOrIp, member, driverType, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, privateMappingPolicy, publicMappingPolicy, mergeData, updateMetadata)
}

func testAccVdiscoverytaskFqdnOrIp(name, fqdnOrIp, username, password, member string, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, mergeData, updateMetadata bool, driverType, privateNetworkViewMappingPolicy, protocol, publicNetworkViewMappingPolicy, selectedRegions string, port int) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_fqdn_or_ip" {
    name                                = %q
    fqdn_or_ip                          = %q
    username                            = %q
    password                            = %q
    member                              = %q
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    merge_data                          = %t
    update_metadata                     = %t
    driver_type                         = %q
    private_network_view_mapping_policy = %q
    protocol                            = %q
    public_network_view_mapping_policy  = %q
    selected_regions                    = %q
    port                                = %d
}
`, name, fqdnOrIp, username, password, member, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, mergeData, updateMetadata, driverType, privateNetworkViewMappingPolicy, protocol, publicNetworkViewMappingPolicy, selectedRegions, port)
}

func testAccVdiscoverytaskGovcloudEnabled(name string, govcloudEnabled bool, username, password, member string, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm bool, driverType, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy string, mergeData, updateMetadata bool, selectedRegions string) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_govcloud_enabled" {
    name                                = %q
    govcloud_enabled                    = %t
    username                            = %q
    password                            = %q
    member                              = %q
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    driver_type                         = %q
    private_network_view_mapping_policy = %q
    public_network_view_mapping_policy  = %q
    merge_data                          = %t
    update_metadata                     = %t
    selected_regions                    = %q
}
`, name, govcloudEnabled, username, password, member, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, driverType, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, mergeData, updateMetadata, selectedRegions)
}

func testAccVdiscoverytaskIdentityVersionV2(name, identityVersion, fqdnOrIp, member, username, password string, useIdentity, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm bool, driverType, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy string, mergeData, updateMetadata bool, port int, protocol string) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_identity_version" {
    name                                = %q
    identity_version                    = %q
    fqdn_or_ip                          = %q
    member                              = %q
    username                            = %q
    password                            = %q
    use_identity                        = %t
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    driver_type                         = %q
    private_network_view_mapping_policy = %q
    public_network_view_mapping_policy  = %q
    merge_data                          = %t
    update_metadata                     = %t
    port                                = %d
    protocol                            = %q
}
`, name, identityVersion, fqdnOrIp, member, username, password, useIdentity, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, driverType, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, mergeData, updateMetadata, port, protocol)
}

func testAccVdiscoverytaskIdentityVersionV3(name, identityVersion, domainName, fqdnOrIp, member, username, password string, useIdentity, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm bool, driverType, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy string, mergeData, updateMetadata bool, port int, protocol string) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_identity_version" {
    name                                = %q
    identity_version                    = %q
    domain_name                         = %q
    fqdn_or_ip                          = %q
    member                              = %q
    username                            = %q
    password                            = %q
    use_identity                        = %t
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    driver_type                         = %q
    private_network_view_mapping_policy = %q
    public_network_view_mapping_policy  = %q
    merge_data                          = %t
    update_metadata                     = %t
    port                                = %d
    protocol                            = %q
}
`, name, identityVersion, domainName, fqdnOrIp, member, username, password, useIdentity, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, driverType, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, mergeData, updateMetadata, port, protocol)
}

func testAccVdiscoverytaskMember(name, member string, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm bool, driverType, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy string, mergeData, updateMetadata bool, selectedRegions, username, password string) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_member" {
    name                                = %q
    member                              = %q
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    driver_type                         = %q
    merge_data                          = %t
    update_metadata                     = %t
    private_network_view_mapping_policy = %q
    public_network_view_mapping_policy  = %q
    selected_regions                    = %q
    username                            = %q
    password                            = %q
}
`, name, member, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, driverType, mergeData, updateMetadata, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, selectedRegions, username, password)
}

func testAccVdiscoverytaskMergeData(name string, mergeData bool, username, password, member string, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm bool, driverType, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, selectedRegions string, updateMetadata bool) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_merge_data" {
    name                                = %q
    merge_data                          = %t
    username                            = %q
    password                            = %q
    member                              = %q
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    driver_type                         = %q
    private_network_view_mapping_policy = %q
    public_network_view_mapping_policy  = %q
    selected_regions                    = %q
    update_metadata                     = %t
}
`, name, mergeData, username, password, member, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, driverType, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, selectedRegions, updateMetadata)
}

func testAccVdiscoverytaskMultipleAccountsSyncPolicyDiscover(name, multipleAccountsSyncPolicy string, syncChildAccounts bool, roleArn, username, password, member string, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm bool, driverType, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy string, mergeData, updateMetadata bool, selectedRegions string) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_multiple_accounts_sync_policy" {
    name                                = %q
    multiple_accounts_sync_policy       = %q
    sync_child_accounts                 = %t
    role_arn                            = %q
    username                            = %q
    password                            = %q
    member                              = %q
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    driver_type                         = %q
    private_network_view_mapping_policy = %q
    public_network_view_mapping_policy  = %q
    merge_data                          = %t
    update_metadata                     = %t
    selected_regions                    = %q
}
`, name, multipleAccountsSyncPolicy, syncChildAccounts, roleArn, username, password, member, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, driverType, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, mergeData, updateMetadata, selectedRegions)
}

func testAccVdiscoverytaskMultipleAccountsSyncPolicyUpload(name, multipleAccountsSyncPolicy string, syncChildAccounts bool, cdiscoveryFile, roleArn, username, password, member string, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm bool, driverType, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy string, mergeData, updateMetadata bool, selectedRegions string) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_multiple_accounts_sync_policy" {
    name                                = %q
    multiple_accounts_sync_policy       = %q
    sync_child_accounts                 = %t
    cdiscovery_file                     = %q
    role_arn                            = %q
    username                            = %q
    password                            = %q
    member                              = %q
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    driver_type                         = %q
    private_network_view_mapping_policy = %q
    public_network_view_mapping_policy  = %q
    merge_data                          = %t
    update_metadata                     = %t
    selected_regions                    = %q
}
`, name, multipleAccountsSyncPolicy, syncChildAccounts, cdiscoveryFile, roleArn, username, password, member, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, driverType, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, mergeData, updateMetadata, selectedRegions)
}

func testAccVdiscoverytaskName(name, member string, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm bool, driverType, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy string, mergeData, updateMetadata bool, selectedRegions, username, password string) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_name" {
    name                                = %q
    member                              = %q
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    driver_type                         = %q
    merge_data                          = %t
    update_metadata                     = %t
    private_network_view_mapping_policy = %q
    public_network_view_mapping_policy  = %q
    selected_regions                    = %q
    username                            = %q
    password                            = %q
}
`, name, member, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, driverType, mergeData, updateMetadata, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, selectedRegions, username, password)
}

func testAccVdiscoverytaskNetworkFilter(name string, enableFilter bool, networkFilter string, networkList []string, username, password, member string, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm bool, driverType, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy string, mergeData, updateMetadata bool, selectedRegions string) string {
	networkListHCL := utils.ConvertStringSliceToHCL(networkList)
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_network_filter" {
    name                                = %q
    enable_filter                       = %t
    network_filter                      = %q
    network_list                        = %s
    username                            = %q
    password                            = %q
    member                              = %q
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    driver_type                         = %q
    private_network_view_mapping_policy = %q
    public_network_view_mapping_policy  = %q
    merge_data                          = %t
    update_metadata                     = %t
    selected_regions                    = %q
}
`, name, enableFilter, networkFilter, networkListHCL, username, password, member, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, driverType, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, mergeData, updateMetadata, selectedRegions)
}

func testAccVdiscoverytaskNetworkList(name string, networkList []string, username, password, member string, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm bool, driverType, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy string, mergeData, updateMetadata bool, selectedRegions string) string {
	networkListHCL := utils.ConvertStringSliceToHCL(networkList)
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_network_list" {
    name                                = %q
    network_list                        = %s
    username                            = %q
    password                            = %q
    member                              = %q
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    driver_type                         = %q
    private_network_view_mapping_policy = %q
    public_network_view_mapping_policy  = %q
    merge_data                          = %t
    update_metadata                     = %t
    selected_regions                    = %q
}
`, name, networkListHCL, username, password, member, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, driverType, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, mergeData, updateMetadata, selectedRegions)
}

func testAccVdiscoverytaskPassword(name, username, password, member string, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm bool, driverType, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy string, mergeData, updateMetadata bool, selectedRegions string) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_password" {
    name                                = %q
    username                            = %q
    password                            = %q
    member                              = %q
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    driver_type                         = %q
    merge_data                          = %t
    update_metadata                     = %t
    private_network_view_mapping_policy = %q
    public_network_view_mapping_policy  = %q
    selected_regions                    = %q
}
`, name, username, password, member, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, driverType, mergeData, updateMetadata, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, selectedRegions)
}

func testAccVdiscoverytaskPort(name string, port int, username, password, member, fqdnOrIp string, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, mergeData, updateMetadata bool, driverType, privateNetworkViewMappingPolicy, protocol, publicNetworkViewMappingPolicy, selectedRegions string) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_port" {
    name                                = %q
    port                                = %d
    username                            = %q
    password                            = %q
    member                              = %q
    fqdn_or_ip                          = %q
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    merge_data                          = %t
    update_metadata                     = %t
    driver_type                         = %q
    private_network_view_mapping_policy = %q
    protocol                            = %q
    public_network_view_mapping_policy  = %q
    selected_regions                    = %q
}
`, name, port, username, password, member, fqdnOrIp, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, mergeData, updateMetadata, driverType, privateNetworkViewMappingPolicy, protocol, publicNetworkViewMappingPolicy, selectedRegions)
}

func testAccVdiscoverytaskPrivateNetworkView(name, privateNetworkView, privateNetworkViewMappingPolicy, username, password, member string, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm bool, driverType, publicNetworkViewMappingPolicy string, mergeData, updateMetadata bool, selectedRegions string) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_private_network_view" {
    name                                = %q
    private_network_view                = %q
    private_network_view_mapping_policy = %q
    username                            = %q
    password                            = %q
    member                              = %q
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    driver_type                         = %q
    public_network_view_mapping_policy  = %q
    merge_data                          = %t
    update_metadata                     = %t
    selected_regions                    = %q
}
`, name, privateNetworkView, privateNetworkViewMappingPolicy, username, password, member, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, driverType, publicNetworkViewMappingPolicy, mergeData, updateMetadata, selectedRegions)
}

func testAccVdiscoverytaskPrivateNetworkViewMappingPolicyAutoCreate(name, privateNetworkViewMappingPolicy, username, password, member string, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm bool, driverType, publicNetworkViewMappingPolicy string, mergeData, updateMetadata bool, selectedRegions string) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_private_network_view_mapping_policy" {
    name                                = %q
    private_network_view_mapping_policy = %q
    username                            = %q
    password                            = %q
    member                              = %q
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    driver_type                         = %q
    public_network_view_mapping_policy  = %q
    merge_data                          = %t
    update_metadata                     = %t
    selected_regions                    = %q
}
`, name, privateNetworkViewMappingPolicy, username, password, member, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, driverType, publicNetworkViewMappingPolicy, mergeData, updateMetadata, selectedRegions)
}

func testAccVdiscoverytaskPrivateNetworkViewMappingPolicyDirect(name, privateNetworkViewMappingPolicy, privateNetworkView, username, password, member string, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm bool, driverType, publicNetworkViewMappingPolicy string, mergeData, updateMetadata bool, selectedRegions string) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_private_network_view_mapping_policy" {
    name                                = %q
    private_network_view_mapping_policy = %q
    private_network_view                = %q
    username                            = %q
    password                            = %q
    member                              = %q
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    driver_type                         = %q
    public_network_view_mapping_policy  = %q
    merge_data                          = %t
    update_metadata                     = %t
    selected_regions                    = %q
}
`, name, privateNetworkViewMappingPolicy, privateNetworkView, username, password, member, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, driverType, publicNetworkViewMappingPolicy, mergeData, updateMetadata, selectedRegions)
}

func testAccVdiscoverytaskProtocol(name, protocol, username, password, member, fqdnOrIp string, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, mergeData, updateMetadata bool, driverType, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, selectedRegions string, port int) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_protocol" {
    name                                = %q
    protocol                            = %q
    username                            = %q
    password                            = %q
    member                              = %q
    fqdn_or_ip                          = %q
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    merge_data                          = %t
    update_metadata                     = %t
    driver_type                         = %q
    private_network_view_mapping_policy = %q
    public_network_view_mapping_policy  = %q
    selected_regions                    = %q
    port                                = %d
}
`, name, protocol, username, password, member, fqdnOrIp, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, mergeData, updateMetadata, driverType, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, selectedRegions, port)
}

func testAccVdiscoverytaskPublicNetworkView(name, publicNetworkView, publicNetworkViewMappingPolicy, username, password, member string, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm bool, driverType, privateNetworkViewMappingPolicy string, mergeData, updateMetadata bool, selectedRegions string) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_public_network_view" {
    name                                = %q
    public_network_view                 = %q
    public_network_view_mapping_policy  = %q
    username                            = %q
    password                            = %q
    member                              = %q
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    driver_type                         = %q
    private_network_view_mapping_policy = %q
    merge_data                          = %t
    update_metadata                     = %t
    selected_regions                    = %q
}
`, name, publicNetworkView, publicNetworkViewMappingPolicy, username, password, member, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, driverType, privateNetworkViewMappingPolicy, mergeData, updateMetadata, selectedRegions)
}

func testAccVdiscoverytaskPublicNetworkViewMappingPolicyAutoCreate(name, publicNetworkViewMappingPolicy, username, password, member string, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm bool, driverType, privateNetworkViewMappingPolicy string, mergeData, updateMetadata bool, selectedRegions string) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_public_network_view_mapping_policy" {
    name                                = %q
    public_network_view_mapping_policy  = %q
    username                            = %q
    password                            = %q
    member                              = %q
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    driver_type                         = %q
    private_network_view_mapping_policy = %q
    merge_data                          = %t
    update_metadata                     = %t
    selected_regions                    = %q
}
`, name, publicNetworkViewMappingPolicy, username, password, member, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, driverType, privateNetworkViewMappingPolicy, mergeData, updateMetadata, selectedRegions)
}

func testAccVdiscoverytaskPublicNetworkViewMappingPolicyDirect(name, publicNetworkViewMappingPolicy, publicNetworkView, username, password, member string, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm bool, driverType, privateNetworkViewMappingPolicy string, mergeData, updateMetadata bool, selectedRegions string) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_public_network_view_mapping_policy" {
    name                                = %q
    public_network_view_mapping_policy  = %q
    public_network_view                 = %q
    username                            = %q
    password                            = %q
    member                              = %q
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    driver_type                         = %q
    private_network_view_mapping_policy = %q
    merge_data                          = %t
    update_metadata                     = %t
    selected_regions                    = %q
}
`, name, publicNetworkViewMappingPolicy, publicNetworkView, username, password, member, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, driverType, privateNetworkViewMappingPolicy, mergeData, updateMetadata, selectedRegions)
}

func testAccVdiscoverytaskRoleArn(name, roleArn, multipleAccountsSyncPolicy string, syncChildAccounts bool, username, password, member string, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm bool, driverType, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy string, mergeData, updateMetadata bool, selectedRegions string) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_role_arn" {
    name                                = %q
    role_arn                            = %q
    multiple_accounts_sync_policy       = %q
    sync_child_accounts                 = %t
    username                            = %q
    password                            = %q
    member                              = %q
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    driver_type                         = %q
    private_network_view_mapping_policy = %q
    public_network_view_mapping_policy  = %q
    merge_data                          = %t
    update_metadata                     = %t
    selected_regions                    = %q
}
`, name, roleArn, multipleAccountsSyncPolicy, syncChildAccounts, username, password, member, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, driverType, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, mergeData, updateMetadata, selectedRegions)
}

func testAccVdiscoverytaskScheduledRun(name string, scheduledRun map[string]any, username, password, fqdnOrIp, member, driverType string, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm bool, privateMappingPolicy, publicMappingPolicy string, mergeData, updateMetadata bool) string {
	scheduledRunHCL := utils.ConvertMapToHCL(scheduledRun)
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_scheduled_run_block_azure" {
    name                                = %q
    scheduled_run                       = %s
    username                            = %q
    password                            = %q
    fqdn_or_ip                          = %q
    member                              = %q
    driver_type                         = %q
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    private_network_view_mapping_policy = %q
    public_network_view_mapping_policy  = %q
    merge_data                          = %t
    update_metadata                     = %t
}
`, name, scheduledRunHCL, username, password, fqdnOrIp, member, driverType, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, privateMappingPolicy, publicMappingPolicy, mergeData, updateMetadata)
}

func testAccVdiscoverytaskSelectedRegions(name, selectedRegions, username, password, member string, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm bool, driverType, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy string, mergeData, updateMetadata bool) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_selected_regions" {
    name                                = %q
    selected_regions                    = %q
    username                            = %q
    password                            = %q
    member                              = %q
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    driver_type                         = %q
    private_network_view_mapping_policy = %q
    public_network_view_mapping_policy  = %q
    merge_data                          = %t
    update_metadata                     = %t
}
`, name, selectedRegions, username, password, member, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, driverType, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, mergeData, updateMetadata)
}

func testAccVdiscoverytaskServiceAccountFile(name, driverType, member, serviceAccountFile, cdiscoveryFile, multipleAccountsSyncPolicy string, mergeData, updateMetadata, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm bool, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy string, enabled bool) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_gcp_service_account_file" {
    name                                = %q
    driver_type                         = %q
    member                              = %q
    service_account_file                = %q
    cdiscovery_file                     = %q
    multiple_accounts_sync_policy       = %q
    merge_data                          = %t
    update_metadata                     = %t
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    private_network_view_mapping_policy = %q
    public_network_view_mapping_policy  = %q
    enabled                             = %t
}
`, name, driverType, member, serviceAccountFile, cdiscoveryFile, multipleAccountsSyncPolicy, mergeData, updateMetadata, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, enabled)
}

func testAccVdiscoverytaskSyncChildAccounts(name string, syncChildAccounts bool, roleArn, multipleAccountsSyncPolicy, username, password, member string, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm bool, driverType, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy string, mergeData, updateMetadata bool, selectedRegions string) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_sync_child_accounts" {
    name                                = %q
    sync_child_accounts                 = %t
    role_arn                            = %q
    multiple_accounts_sync_policy       = %q
    username                            = %q
    password                            = %q
    member                              = %q
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    driver_type                         = %q
    private_network_view_mapping_policy = %q
    public_network_view_mapping_policy  = %q
    merge_data                          = %t
    update_metadata                     = %t
    selected_regions                    = %q
}
`, name, syncChildAccounts, roleArn, multipleAccountsSyncPolicy, username, password, member, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, driverType, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, mergeData, updateMetadata, selectedRegions)
}

func testAccVdiscoverytaskUpdateDnsViewPrivateIpFalse(name string, updateDnsViewPrivateIp bool, username, password, member string, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm bool, driverType, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy string, mergeData, updateMetadata bool, selectedRegions string) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_update_dns_view_private_ip" {
    name                                = %q
    update_dns_view_private_ip          = %t
    username                            = %q
    password                            = %q
    member                              = %q
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    driver_type                         = %q
    private_network_view_mapping_policy = %q
    public_network_view_mapping_policy  = %q
    merge_data                          = %t
    update_metadata                     = %t
    selected_regions                    = %q
}
`, name, updateDnsViewPrivateIp, username, password, member, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, driverType, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, mergeData, updateMetadata, selectedRegions)
}

func testAccVdiscoverytaskUpdateDnsViewPrivateIpTrue(name string, updateDnsViewPrivateIp bool, dnsViewPrivateIp, username, password, member string, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm bool, driverType, privateNetworkViewMappingPolicy, privateNetworkView, publicNetworkViewMappingPolicy string, mergeData, updateMetadata bool, selectedRegions string) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_update_dns_view_private_ip" {
    name                                = %q
    update_dns_view_private_ip          = %t
    dns_view_private_ip                 = %q
    username                            = %q
    password                            = %q
    member                              = %q
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    driver_type                         = %q
    private_network_view_mapping_policy = %q
    private_network_view                = %q
    public_network_view_mapping_policy  = %q
    merge_data                          = %t
    update_metadata                     = %t
    selected_regions                    = %q
}
`, name, updateDnsViewPrivateIp, dnsViewPrivateIp, username, password, member, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, driverType, privateNetworkViewMappingPolicy, privateNetworkView, publicNetworkViewMappingPolicy, mergeData, updateMetadata, selectedRegions)
}

func testAccVdiscoverytaskUpdateDnsViewPublicIpFalse(name string, updateDnsViewPublicIp bool, username, password, member string, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm bool, driverType, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy string, mergeData, updateMetadata bool, selectedRegions string) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_update_dns_view_public_ip" {
    name                                = %q
    update_dns_view_public_ip           = %t
    username                            = %q
    password                            = %q
    member                              = %q
    driver_type                         = %q
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    private_network_view_mapping_policy = %q
    public_network_view_mapping_policy  = %q
    merge_data                          = %t
    update_metadata                     = %t
    selected_regions                    = %q
}
`, name, updateDnsViewPublicIp, username, password, member, driverType, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, mergeData, updateMetadata, selectedRegions)
}

func testAccVdiscoverytaskUpdateDnsViewPublicIpTrue(name string, updateDnsViewPublicIp bool, dnsViewPublicIp, username, password, member string, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm bool, driverType, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, publicNetworkView string, mergeData, updateMetadata bool, selectedRegions string) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_update_dns_view_public_ip" {
    name                                = %q
    update_dns_view_public_ip           = %t
    dns_view_public_ip                  = %q
    username                            = %q
    password                            = %q
    member                              = %q
    driver_type                         = %q
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    private_network_view_mapping_policy = %q
    public_network_view_mapping_policy  = %q
    public_network_view                 = %q
    merge_data                          = %t
    update_metadata                     = %t
    selected_regions                    = %q
}
`, name, updateDnsViewPublicIp, dnsViewPublicIp, username, password, member, driverType, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, publicNetworkView, mergeData, updateMetadata, selectedRegions)
}

func testAccVdiscoverytaskUpdateMetadata(name string, updateMetadata bool, username, password, member string, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm bool, driverType, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, selectedRegions string, mergeData bool) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_update_metadata" {
    name                                = %q
    update_metadata                     = %t
    username                            = %q
    password                            = %q
    member                              = %q
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    driver_type                         = %q
    private_network_view_mapping_policy = %q
    public_network_view_mapping_policy  = %q
    selected_regions                    = %q
    merge_data                          = %t
}
`, name, updateMetadata, username, password, member, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, driverType, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, selectedRegions, mergeData)
}

func testAccVdiscoverytaskUseIdentity(name string, useIdentity bool, fqdnOrIp string, port int, protocol, member, identityVersion, username, password string, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm bool, driverType, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy string, mergeData, updateMetadata bool) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_use_identity" {
    name                                = %q
    use_identity                        = %t
    fqdn_or_ip                          = %q
    port                                = %d
    protocol                            = %q
    member                              = %q
    identity_version                    = %q
    username                            = %q
    password                            = %q
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    driver_type                         = %q
    merge_data                          = %t
    private_network_view_mapping_policy = %q
    public_network_view_mapping_policy  = %q
    update_metadata                     = %t
}
`, name, useIdentity, fqdnOrIp, port, protocol, member, identityVersion, username, password, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, driverType, mergeData, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, updateMetadata)
}

func testAccVdiscoverytaskUsername(name, username, password, member string, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm bool, driverType, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy string, mergeData, updateMetadata bool, selectedRegions string) string {
	return fmt.Sprintf(`
resource "nios_discovery_vdiscovery_task" "test_username" {
    name                                = %q
    username                            = %q
    password                            = %q
    member                              = %q
    auto_consolidate_cloud_ea           = %t
    auto_consolidate_managed_tenant     = %t
    auto_consolidate_managed_vm         = %t
    driver_type                         = %q
    merge_data                          = %t
    update_metadata                     = %t
    private_network_view_mapping_policy = %q
    public_network_view_mapping_policy  = %q
    selected_regions                    = %q
}
`, name, username, password, member, autoConsolidateCloudEa, autoConsolidateManagedTenant, autoConsolidateManagedVm, driverType, mergeData, updateMetadata, privateNetworkViewMappingPolicy, publicNetworkViewMappingPolicy, selectedRegions)
}

// Helper function to get test data path
func getTestDataPath() string {
	_, filename, _, _ := runtime.Caller(0)
	testDir := filepath.Dir(filename)
	return filepath.Join(testDir, "..", "..", "testdata", "nios_discovery_vdiscovery_task")
}
