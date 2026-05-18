provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

module "node1" {
  source = "github.com/infobloxopen/terraform-provider-nios//modules/nios_deploy_oci?ref=nios_v9.1.0"

  default_admin_password = var.default_admin_password
  remote_console_enabled = var.remote_console_enabled
  nios_license           = var.nios_license
  compartment_id           = var.compartment_id
  availability_domain      = var.availability_domain
  bucket_name              = var.bucket_name
  create_bucket            = var.create_bucket
  nios_qcow2_local_path    = var.nios_qcow2_local_path
  nios_object_name         = var.nios_object_name
  image_name               = var.image_name
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
}
