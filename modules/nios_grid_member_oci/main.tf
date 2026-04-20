// Object Storage namespace
data "oci_objectstorage_namespace" "ns" {
  compartment_id = var.compartment_id
}

// Object Storage Bucket
resource "oci_objectstorage_bucket" "nios_bucket" {
  count = var.create_bucket ? 1 : 0

  compartment_id = var.compartment_id
  namespace      = data.oci_objectstorage_namespace.ns.namespace
  name           = var.bucket_name
  access_type    = "NoPublicAccess"
}

locals {
  bucket_name = var.create_bucket ? oci_objectstorage_bucket.nios_bucket[0].name : var.bucket_name

  // NIOS model -> VM.Standard3.Flex shape config
  nios_model_config = {
    "IB-V926"  = { ocpus = 4, memory = 32 }
    "IB-V1516" = { ocpus = 6, memory = 64 }
    "IB-V1526" = { ocpus = 8, memory = 64 }
    "IB-V2326" = { ocpus = 10, memory = 192 }
    "IB-V4126" = { ocpus = 16, memory = 384 }
    "IB-V5005" = { ocpus = var.instance_ocpus, memory = var.instance_memory_in_gbs }
  }

  effective_ocpus  = local.nios_model_config[var.nios_model].ocpus
  effective_memory = local.nios_model_config[var.nios_model].memory

  // Cloud-Init resolution
  cloud_init_user_data = (
    var.cloud_init_content != "" ? base64encode(var.cloud_init_content) :
    var.cloud_init_script_path != "" ? filebase64(var.cloud_init_script_path) :
    base64encode(templatefile("${path.module}/user_data.tftpl", {
      nios_license           = var.nios_license
      remote_console_enabled = var.remote_console_enabled ? "y" : "n"
      default_admin_password = var.default_admin_password
    }))
  )
}

// Upload NIOS QCOW2 to Object Storage
resource "oci_objectstorage_object" "nios_qcow2" {
  namespace    = data.oci_objectstorage_namespace.ns.namespace
  bucket       = var.bucket_name
  object       = var.nios_object_name
  source       = var.nios_qcow2_local_path
  content_type = "application/octet-stream"

  depends_on = [oci_objectstorage_bucket.nios_bucket]
}

// Custom Image — Import from Object Storage
// Both instances share one image. Import can take 30–60 minutes.
resource "oci_core_image" "nios_image" {
  compartment_id = var.compartment_id
  display_name   = var.image_name

  image_source_details {
    source_type       = "objectStorageTuple"
    namespace_name    = data.oci_objectstorage_namespace.ns.namespace
    bucket_name       = var.bucket_name
    object_name       = var.nios_object_name
    operating_system  = "Linux"
    source_image_type = "QCOW2"
  }

  launch_mode = "PARAVIRTUALIZED"

  depends_on = [oci_objectstorage_object.nios_qcow2]

  timeouts {
    create = "60m"
  }
}

// Compute Instance
resource "oci_core_instance" "nios_instance" {

  compartment_id      = var.compartment_id
  display_name        = var.instance_name
  availability_domain = var.availability_domain

  // Boot image
  source_details {
    source_type = "image"
    source_id   = oci_core_image.nios_image.id
  }

  // Shape
  shape = var.nios_version_gte_902 ? "VM.Standard3.Flex" : var.legacy_shape

  dynamic "shape_config" {
    for_each = var.nios_version_gte_902 ? [1] : []
    content {
      ocpus         = local.effective_ocpus
      memory_in_gbs = local.effective_memory
    }
  }

  // Primary VNIC — MGMT (eth0)
  create_vnic_details {
    display_name           = var.mgmt_vnic_name
    subnet_id              = var.mgmt_subnet_id
    assign_public_ip       = var.mgmt_assign_public_ip
    skip_source_dest_check = false
  }

  // Cloud-Init
  metadata = local.cloud_init_user_data != "" ? {
    user_data = local.cloud_init_user_data
  } : {}

  lifecycle {
    ignore_changes = [metadata]
  }
}

// Secondary VNICs — LAN1 (one per instance)
resource "oci_core_vnic_attachment" "lan1_vnic_attachment" {

  instance_id  = oci_core_instance.nios_instance.id
  display_name = "${var.instance_name}-lan1-attachment"

  create_vnic_details {
    display_name           = var.lan1_vnic_name
    subnet_id              = var.lan1_subnet_id
    assign_public_ip       = var.lan1_assign_public_ip
    skip_source_dest_check = false
  }
}

// Reporting Block Volume (optional — member instance only)
resource "oci_core_volume" "reporting_volume" {
  count = var.enable_reporting_volume ? 1 : 0

  compartment_id      = var.compartment_id
  availability_domain = var.availability_domain
  display_name        = var.reporting_volume_name
  size_in_gbs         = var.reporting_volume_size_gb
}

resource "oci_core_volume_attachment" "reporting_volume_attachment" {
  count = var.enable_reporting_volume ? 1 : 0

  attachment_type = "paravirtualized"
  // Attach reporting volume to the member instance [1]
  instance_id  = oci_core_instance.nios_instance.id
  volume_id    = oci_core_volume.reporting_volume[0].id
  is_read_only = false
  is_shareable = false
}
