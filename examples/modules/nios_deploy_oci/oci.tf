provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

module "node1" {
  source = "github.com/infobloxopen/terraform-provider-nios//modules/nios_deploy_oci?ref=nios_v9.1.0"

  default_admin_password   = var.default_admin_password
  remote_console_enabled   = var.remote_console_enabled
  nios_license             = var.nios_license
  image_id                 = var.image_id
  compartment_id           = var.compartment_id
  availability_domain      = var.availability_domain
  instance_name            = var.instance_name
  nios_model               = var.nios_model
  nios_version_gte_9xx     = var.nios_version_gte_9xx
  legacy_shape             = var.legacy_shape
  instance_ocpus           = var.instance_ocpus
  instance_memory_in_gbs   = var.instance_memory_in_gbs
  mgmt_subnet_id           = var.mgmt_subnet_id
  mgmt_assign_public_ip    = var.mgmt_assign_public_ip
  lan1_subnet_id           = var.lan1_subnet_id
  lan1_assign_public_ip    = var.lan1_assign_public_ip
  enable_reporting_volume  = var.enable_reporting_volume
  reporting_volume_name    = var.reporting_volume_name
  reporting_volume_size_gb = var.reporting_volume_size_gb
  freeform_tags            = var.freeform_tags
}

// HA Configuration
// To create primary node for HA pair, set enable_ha to true and is_primary to true. 
module "node2" {
  source = "github.com/infobloxopen/terraform-provider-nios//modules/nios_deploy_oci?ref=nios_v9.1.0"

  default_admin_password   = var.default_admin_password
  remote_console_enabled   = var.remote_console_enabled
  nios_license             = var.nios_license
  image_id                 = var.image_id
  compartment_id           = var.compartment_id
  availability_domain      = var.availability_domain
  instance_name            = "${var.instance_name}-1"
  nios_model               = var.nios_model
  nios_version_gte_9xx     = var.nios_version_gte_9xx
  legacy_shape             = var.legacy_shape
  instance_ocpus           = var.instance_ocpus
  instance_memory_in_gbs   = var.instance_memory_in_gbs
  mgmt_subnet_id           = var.mgmt_subnet_id
  mgmt_assign_public_ip    = var.mgmt_assign_public_ip
  freeform_tags            = var.freeform_tags
  lan1_subnet_id           = var.lan1_subnet_id
  lan1_assign_public_ip    = var.lan1_assign_public_ip
  enable_reporting_volume  = var.enable_reporting_volume
  reporting_volume_name    = var.reporting_volume_name
  reporting_volume_size_gb = var.reporting_volume_size_gb

  enable_ha    = true
  ha_subnet_id = var.ha_subnet_id
  is_primary   = true
}

// To create secondary node for HA pair, set enable_ha to true and is_primary to false.
module "node3" {
  source = "github.com/infobloxopen/terraform-provider-nios//modules/nios_deploy_oci?ref=nios_v9.1.0"

  default_admin_password   = var.default_admin_password
  remote_console_enabled   = var.remote_console_enabled
  nios_license             = var.nios_license
  image_id                 = var.image_id
  compartment_id           = var.compartment_id
  availability_domain      = var.availability_domain
  instance_name            = "${var.instance_name}-2"
  nios_model               = var.nios_model
  nios_version_gte_9xx     = var.nios_version_gte_9xx
  legacy_shape             = var.legacy_shape
  instance_ocpus           = var.instance_ocpus
  instance_memory_in_gbs   = var.instance_memory_in_gbs
  mgmt_subnet_id           = var.mgmt_subnet_id
  mgmt_assign_public_ip    = var.mgmt_assign_public_ip
  freeform_tags            = var.freeform_tags
  lan1_subnet_id           = var.lan1_subnet_id
  lan1_assign_public_ip    = var.lan1_assign_public_ip
  enable_reporting_volume  = var.enable_reporting_volume
  reporting_volume_name    = var.reporting_volume_name
  reporting_volume_size_gb = var.reporting_volume_size_gb

  enable_ha    = true
  ha_subnet_id = var.ha_subnet_id
  is_primary   = false
}
