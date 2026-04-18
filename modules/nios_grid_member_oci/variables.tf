# OCI Provider Authentication
variable "tenancy_ocid" {
  description = "OCID of your OCI tenancy."
  type        = string
}

variable "user_ocid" {
  description = "OCID of the OCI IAM user used for API authentication."
  type        = string
}

variable "fingerprint" {
  description = "Fingerprint of the API signing key."
  type        = string
}

variable "private_key_path" {
  description = "Absolute local path to the OCI API private key (PEM file)."
  type        = string
}

variable "region" {
  description = "OCI region identifier (e.g. us-ashburn-1)."
  type        = string
}

# Compartment
variable "compartment_id" {
  description = "OCID of the compartment in which all resources will be created."
  type        = string
}

# Object Storage — Bucket
variable "create_bucket" {
  description = "Set to true to create a new bucket; false to reuse an existing one."
  type        = bool
  default     = true
}

variable "bucket_name" {
  description = "Name of the Object Storage bucket for the NIOS QCOW2 image."
  type        = string
}

# Object Storage — QCOW2 Upload
variable "nios_qcow2_local_path" {
  description = "Absolute local path to the NIOS QCOW2 image file."
  type        = string
}

variable "nios_object_name" {
  description = "Object name to store the QCOW2 as in the bucket."
  type        = string
}

# Custom Image
variable "image_name" {
  description = "Display name for the custom OCI image imported from the QCOW2."
  type        = string
  default     = "nios-custom-image"
}

# Compute Instances
variable "instance_name" {
  description = "Base display name. Instances will be named <name>-gm and <name>-member."
  type        = string
  default     = "nios"
}

variable "availability_domain" {
  description = "Full availability domain name (e.g. Uocm:US-ASHBURN-AD-1)."
  type        = string
}

# Instance Shape
variable "nios_version_gte_902" {
  description = "true → VM.Standard3.Flex (NIOS >= 9.0.2). false → legacy_shape."
  type        = bool
  default     = true
}

variable "nios_model" {
  description = <<-EOT
    NIOS appliance model — sets OCPUs and memory for Flex shape.
    One of: IB-V926, IB-V1516, IB-V1526, IB-V2326, IB-V4126, IB-V5005.
  EOT
  type    = string
  default = "IB-V926"
  validation {
    condition     = contains(["IB-V926", "IB-V1516", "IB-V1526", "IB-V2326", "IB-V4126", "IB-V5005"], var.nios_model)
    error_message = "nios_model must be one of: IB-V926, IB-V1516, IB-V1526, IB-V2326, IB-V4126, IB-V5005."
  }
}

variable "instance_ocpus" {
  description = "OCPUs — used only for IB-V5005."
  type        = number
  default     = 4
}

variable "instance_memory_in_gbs" {
  description = "Memory in GB — used only for IB-V5005."
  type        = number
  default     = 32
}

variable "legacy_shape" {
  description = "Fixed shape for NIOS < 9.0.2 (e.g. VM.Standard2.2)."
  type        = string
  default     = "VM.Standard2.2"
}

# Primary VNIC — MGMT (eth0)
variable "mgmt_vnic_name" {
  description = "Display name for the primary (MGMT) VNIC."
  type        = string
  default     = "nios-mgmt-vnic"
}

variable "mgmt_subnet_id" {
  description = "OCID of the subnet for the MGMT interface."
  type        = string
}

variable "mgmt_assign_public_ip" {
  description = "Assign a public IP to the MGMT VNIC."
  type        = bool
  default     = false
}

# Secondary VNIC — LAN1
variable "lan1_vnic_name" {
  description = "Display name for the secondary (LAN1) VNIC."
  type        = string
  default     = "nios-lan1-vnic"
}

variable "lan1_subnet_id" {
  description = "OCID of the subnet for the LAN1 interface."
  type        = string
}

variable "lan1_assign_public_ip" {
  description = "Assign a public IP to the LAN1 VNIC."
  type        = bool
  default     = false
}

# Cloud-Init
variable "cloud_init_script_path" {
  description = "Path to a cloud-init YAML file. Leave empty to skip."
  type        = string
  default     = ""
}

variable "cloud_init_content" {
  description = "Inline cloud-init YAML. Takes precedence over cloud_init_script_path."
  type        = string
  default     = ""
}

# Reporting Block Volume (optional — attached to member)
variable "enable_reporting_volume" {
  description = "Create and attach a reporting block volume to the Grid Member."
  type        = bool
  default     = false
}

variable "reporting_volume_name" {
  description = "Display name for the reporting block volume."
  type        = string
  default     = "nios-reporting-volume"
}

variable "reporting_volume_size_gb" {
  description = "Size in GB for the reporting volume. Minimum 250 GB recommended."
  type        = number
  default     = 250
}

variable "nios_license" {
  description = "NIOS temporary license string."
  type        = string
  default     = "nios IB-V825 enterprise dns dhcp cloud"
}

variable "remote_console_enabled" {
  description = "Enable remote console access."
  type        = bool
  default     = true
}

variable "default_admin_password" {
  description = "Default admin password for NIOS."
  type        = string
  sensitive   = true
}
