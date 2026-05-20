// Retrieve information about existing Mgmt subnetwork
data "google_compute_subnetwork" "mgmt" {
  name    = var.mgmt_subnet_name
  region  = var.region
  project = var.project
}

// Retrieve information about existing LAN subnetwork
data "google_compute_subnetwork" "lan1" {
  name    = var.lan1_subnet_name
  region  = var.region
  project = var.project
}

// Retrieve information about existing HA subnetwork
data "google_compute_subnetwork" "ha" {
  count   = var.enable_ha ? 1 : 0
  name    = var.ha_subnet_name
  region  = var.region
  project = var.project
}

locals {
  // Machine-type lookup: NIOS model -> GCP machine type
  machine_type_map = {
    "IB-V825"  = "n2-standard-2"
    "IB-V1425" = "n2-standard-4"
    "IB-V2225" = "n2-standard-8"
    "IB-V4025" = "n2-standard-16"
    "TE-V810"  = "n2-standard-2"
    "TE-V1410" = "n2-standard-4"
    "TE-V2210" = "n2-standard-8"
    "TE-V4010" = "n2-standard-16"
    "CP-V800"  = "n2-standard-2"
    "CP-V1400" = "n2-standard-4"
    "CP-V2200" = "n2-standard-8"
    "CP-V4000" = "n2-standard-16"
  }

  // Image self-link
  image = "projects/${var.project}/global/images/${var.image_name}"

  // Subnetwork self-links
  subnetwork_mgmt = "projects/${var.project}/regions/${var.region}/subnetworks/${var.mgmt_subnet_name}"
  subnetwork_lan1 = "projects/${var.project}/regions/${var.region}/subnetworks/${var.lan1_subnet_name}"
  subnetwork_ha   = var.ha_subnet_name != null ? "projects/${var.project}/regions/${var.region}/subnetworks/${var.ha_subnet_name}" : null
}

// Manage a Google Compute Instance for NIOS Grid Member
resource "google_compute_instance" "grid" {
  name         = var.name
  machine_type = try(local.machine_type_map[var.nios_model], var.machine_type)
  zone         = var.zone

  labels = var.labels

  boot_disk {
    initialize_params {
      image  = local.image
      type   = var.boot_disk_type
      size   = var.boot_disk_size
      labels = var.labels
    }
  }

  // nic0 – MGMT
  network_interface {
    subnetwork = local.subnetwork_mgmt
    stack_type = var.enable_ipv6 ? "IPV4_IPV6" : "IPV4_ONLY"
  }

  // nic1 – LAN1
  network_interface {
    subnetwork = local.subnetwork_lan1
    stack_type = var.enable_ipv6 ? "IPV4_IPV6" : "IPV4_ONLY"
  }

  // nic2 – HA (optional)
  dynamic "network_interface" {
    for_each = var.enable_ha ? [1] : []
    content {
      subnetwork = local.subnetwork_ha
      dynamic "alias_ip_range" {
        for_each = var.is_primary ? [1] : []
        content {
          ip_cidr_range = "/32"
        }
      }
    }
  }

  metadata = {
    user-data = templatefile("${path.module}/user_data.tftpl", {
      nios_license           = var.nios_license
      remote_console_enabled = var.remote_console_enabled ? "y" : "n"
      default_admin_password = var.default_admin_password
    })
  }

  dynamic "service_account" {
    for_each = var.service_account_email != null ? [1] : []
    content {
      email  = var.service_account_email
      scopes = var.service_account_scopes
    }
  }
}
