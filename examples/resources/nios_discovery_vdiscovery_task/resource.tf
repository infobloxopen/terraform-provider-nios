// Create an AWS Vdiscovery Task
resource "nios_discovery_vdiscovery_task" "aws_comprehensive" {
  name                                = "AWS-vdiscoverytask"
  driver_type                         = "AWS"
  member                              = "infoblox.localdomain"
  merge_data                          = true
  update_metadata                     = true
  auto_consolidate_cloud_ea           = true
  auto_consolidate_managed_tenant     = false
  auto_consolidate_managed_vm         = true
  private_network_view_mapping_policy = "DIRECT"
  private_network_view                = "default"
  public_network_view_mapping_policy  = "DIRECT"
  public_network_view                 = "default"
  selected_regions                    = "us-east-1"
  sync_child_accounts                 = true
  multiple_accounts_sync_policy       = "DISCOVER"
  role_arn                            = "arn:aws:iam::123456789012:role/InfobloxDiscoveryRole"
  username                            = "aws-access-key-id"
  password                            = "aws-secret-access-key"
  update_dns_view_private_ip          = true
  dns_view_private_ip                 = "default"
  update_dns_view_public_ip           = true
  dns_view_public_ip                  = "default"
  auto_create_dns_record              = true
  auto_create_dns_hostname_template   = "$${vm_name}.domain.com"
  auto_create_dns_record_type         = "A_PTR_RECORD"
  enable_filter                       = true
  network_filter                      = "INCLUDE"
  network_list                        = ["10.0.0.0/8", "172.16.0.0/12", "192.168.0.0/16"]
  comment                             = "Example AWS Vdiscovery Task"
  enabled                             = true
}

// Create an Azure Vdiscovery Task
resource "nios_discovery_vdiscovery_task" "azure_vdiscoverytask" {
  name                                = "AZURE-vdiscoverytask"
  driver_type                         = "AZURE"
  member                              = "infoblox.localdomain"
  username                            = "azure_client_id"
  password                            = "azure_client_secret"
  fqdn_or_ip                          = "tenant_id"
  auto_consolidate_cloud_ea           = true
  auto_consolidate_managed_tenant     = true
  auto_consolidate_managed_vm         = true
  merge_data                          = true
  update_metadata                     = false
  private_network_view_mapping_policy = "AUTO_CREATE"
  public_network_view_mapping_policy  = "AUTO_CREATE"
}

// Create a VMware Vdiscovery Task
resource "nios_discovery_vdiscovery_task" "vmware_vdiscoverytask" {
  name                                = "VMWARE-vdiscoverytask"
  driver_type                         = "VMWARE"
  member                              = "infoblox.localdomain"
  fqdn_or_ip                          = "10.0.0.0"
  username                            = "vc_admin"
  password                            = "vmware_password"
  protocol                            = "HTTPS"
  port                                = 443
  allow_unsecured_connection          = true
  auto_consolidate_cloud_ea           = true
  auto_consolidate_managed_tenant     = true
  auto_consolidate_managed_vm         = true
  merge_data                          = true
  update_metadata                     = false
  private_network_view_mapping_policy = "AUTO_CREATE"
  public_network_view_mapping_policy  = "AUTO_CREATE"
}

//Create an OpenStack Vdiscovery Task
resource "nios_discovery_vdiscovery_task" "openstack_vdiscoverytask" {
  name                                = "OPENSTACK-vdiscoverytask"
  driver_type                         = "OPENSTACK"
  member                              = "infoblox.localdomain"
  fqdn_or_ip                          = "10.15.0.0"
  username                            = "openstack_user"
  password                            = "openstack_password"
  protocol                            = "HTTP"
  port                                = 80
  identity_version                    = "KEYSTONE_V2"
  use_identity                        = true
  auto_consolidate_cloud_ea           = true
  auto_consolidate_managed_tenant     = true
  auto_consolidate_managed_vm         = true
  merge_data                          = true
  update_metadata                     = false
  private_network_view_mapping_policy = "AUTO_CREATE"
  public_network_view_mapping_policy  = "AUTO_CREATE"
}

//Create a GCP Vdiscovery Task
resource "nios_discovery_vdiscovery_task" "gcp_vdiscoverytask" {
  name                                = "GCP-vdiscoverytask"
  driver_type                         = "GCP"
  member                              = "infoblox.localdomain"
  service_account_file                = "</path/to/your/service_account_file.json>"
  cdiscovery_file                     = "</path/to/your/cdiscovery_file.json>"
  multiple_accounts_sync_policy       = "UPLOAD"
  sync_child_accounts                 = true
  merge_data                          = false
  update_metadata                     = true
  auto_consolidate_cloud_ea           = true
  auto_consolidate_managed_tenant     = true
  auto_consolidate_managed_vm         = false
  private_network_view_mapping_policy = "AUTO_CREATE"
  public_network_view_mapping_policy  = "AUTO_CREATE"
}
 