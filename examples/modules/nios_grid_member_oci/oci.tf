provider "oci" {
  tenancy_ocid     = var.tenancy_ocid
  user_ocid        = var.user_ocid
  fingerprint      = var.fingerprint
  private_key_path = var.private_key_path
  region           = var.region
}

module "node1" {
  source = "<path_to_module>"

  tenancy_ocid             = var.tenancy_ocid
  user_ocid                = var.user_ocid
  fingerprint              = var.fingerprint
  private_key_path         = var.private_key_path
  region                   = var.region
  compartment_id           = var.compartment_id
  availability_domain      = var.availability_domain
  bucket_name              = var.bucket_name
  create_bucket            = var.create_bucket
  nios_qcow2_local_path    = var.nios_qcow2_local_path
  nios_object_name         = var.nios_object_name
  image_name               = var.image_name
  instance_name            = var.instance_name
  nios_model               = var.nios_model
  nios_version_gte_902     = var.nios_version_gte_902
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
  cloud_init_content       = var.cloud_init_content
  cloud_init_script_path   = var.cloud_init_script_path
}
