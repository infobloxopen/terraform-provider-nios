// Object Storage namespace (only needed when this module creates the image)
data "oci_objectstorage_namespace" "ns" {
  count          = var.create_image ? 1 : 0
  compartment_id = var.compartment_id
}

// Object Storage Bucket
resource "oci_objectstorage_bucket" "nios_bucket" {
  count = var.create_image && var.create_bucket ? 1 : 0

  compartment_id = var.compartment_id
  namespace      = data.oci_objectstorage_namespace.ns[0].namespace
  name           = var.bucket_name
  access_type    = "NoPublicAccess"
}

locals {
  bucket_name = var.create_image && var.create_bucket ? oci_objectstorage_bucket.nios_bucket[0].name : var.bucket_name

  // Image OCID actually used by the compute instance:
  // - when create_image = true, use the image imported by this module
  // - otherwise, use the existing image OCID provided by the caller
  effective_image_id = var.create_image ? oci_core_image.nios_image[0].id : var.image_id

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
    base64encode(templatefile("${path.module}/user_data.tftpl", {
      nios_license           = var.nios_license
      remote_console_enabled = var.remote_console_enabled ? "y" : "n"
      default_admin_password = var.default_admin_password
    }))
  )
}

// Upload NIOS QCOW2 to Object Storage
resource "oci_objectstorage_object" "nios_qcow2" {
  count = var.create_image ? 1 : 0

  namespace    = data.oci_objectstorage_namespace.ns[0].namespace
  bucket       = local.bucket_name
  object       = var.nios_object_name
  source       = var.nios_qcow2_local_path
  content_type = "application/octet-stream"
}

// Custom Image — Import from Object Storage
// Both instances share one image. Import can take 30–60 minutes.
resource "oci_core_image" "nios_image" {
  count = var.create_image ? 1 : 0

  compartment_id = var.compartment_id
  display_name   = var.image_name

  image_source_details {
    source_type       = "objectStorageTuple"
    namespace_name    = data.oci_objectstorage_namespace.ns[0].namespace
    bucket_name       = local.bucket_name
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
    source_id   = local.effective_image_id
  }

  // Shape
  shape = var.nios_version_gte_9xx ? "VM.Standard3.Flex" : var.legacy_shape

  dynamic "shape_config" {
    for_each = var.nios_version_gte_9xx ? [1] : []
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

  freeform_tags = var.freeform_tags
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

// HA VNIC Attachment (third NIC for HA)
resource "oci_core_vnic_attachment" "ha_vnic_attachment" {
  count        = var.enable_ha ? 1 : 0
  instance_id  = oci_core_instance.nios_instance.id
  display_name = "${var.instance_name}-ha-attachment"

  create_vnic_details {
    display_name           = var.ha_vnic_name
    subnet_id              = var.ha_subnet_id
    assign_public_ip       = var.ha_assign_public_ip
    skip_source_dest_check = true
  }
  depends_on = [oci_core_vnic_attachment.lan1_vnic_attachment]
}

// VIP - Secondary Private IP on HA VNIC (only on primary node)
resource "oci_core_private_ip" "ha_vip" {
  count        = var.enable_ha && var.is_primary ? 1 : 0
  display_name = "${var.instance_name}-vip"
  vnic_id      = oci_core_vnic_attachment.ha_vnic_attachment[0].vnic_id
  lifetime     = "RESERVED"
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
