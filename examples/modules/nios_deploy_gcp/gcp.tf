provider "google" {
  project     = var.project
  region      = var.region
  zone        = var.zone
  credentials = file("<path_to_service_account_key.json>")
}

module "node1" {
  source = "github.com/infobloxopen/terraform-provider-nios//modules/nios_deploy_gcp?ref=nios_v9.1.0"

  project = var.project
  region  = var.region
  zone    = var.zone

  image_name       = var.image_name
  name             = var.name
  nios_model       = var.nios_model
  mgmt_subnet_name = var.mgmt_subnet_name
  lan1_subnet_name = var.lan1_subnet_name

  boot_disk_type = var.boot_disk_type
  boot_disk_size = var.boot_disk_size

  nios_license           = var.nios_license
  default_admin_password = var.default_admin_password

  service_account_email  = var.service_account_email
  service_account_scopes = var.service_account_scopes

  labels = var.labels
}
