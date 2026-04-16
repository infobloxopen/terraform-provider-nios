variable "project_id" {
  description = "GCP project ID where all resources will be deployed."
  type        = string
}

variable "region" {
  description = "GCP region for the deployment."
  type        = string
  default     = "us-west1"
}

variable "zone" {
  description = "GCP zone for the Compute Instance."
  type        = string
  default     = "us-west1-b"
}

variable "image_name" {
  description = "Name of the custom NIOS GCP image in the same project."
  type        = string
}

variable "name" {
  description = "Name for the Compute Instance."
  type        = string
  default     = "nios-gcp-instance"
}

variable "nios_model" {
  description = "NIOS virtual appliance model (e.g. IB-V825, IB-V1425, IB-V2225, IB-V4025, TE-V810, CP-V800). Used for machine type lookup and license."
  type        = string
  default     = "IB-V1425"
}

variable "machine_type" {
  description = "GCP machine type. Used as fallback if nios_model is not found in the machine_type_map."
  type        = string
  default     = "n2-standard-4"
}

variable "mgmt_subnet_name" {
  description = "Name of the management subnetwork (nic0)."
  type        = string
}

variable "lan1_subnet_name" {
  description = "Name of the LAN1 subnetwork (nic1) for grid communication."
  type        = string
}

variable "boot_disk_type" {
  description = "Boot disk type (e.g. pd-standard, pd-ssd, pd-balanced)."
  type        = string
  default     = "pd-standard"
}

variable "boot_disk_size" {
  description = "Boot disk size in GB."
  type        = number
  default     = 250
}

variable "nios_license" {
  description = "NIOS temporary license string (e.g. 'nios IB-V1425 enterprise dns dhcp cloud')."
  type        = string
  default     = "nios IB-V1425 enterprise dns dhcp cloud"
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

variable "service_account_email" {
  description = "Service account email to attach to the instance. Set to null to skip."
  type        = string
  default     = null
}

variable "service_account_scopes" {
  description = "OAuth scopes for the service account."
  type        = list(string)
  default     = ["https://www.googleapis.com/auth/cloud-platform"]
}

variable "labels" {
  description = "Labels to apply to GCP resources."
  type        = map(string)
  default = {
    "product"       = "nios",
    "dontstop"      = "no",
    "dontterminate" = "yes"
  }
}
