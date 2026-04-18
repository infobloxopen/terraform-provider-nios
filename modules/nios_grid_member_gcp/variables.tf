variable "project" {
  description = "The default project to manage resources in."
  type        = string
}

variable "region" {
  description = "The region in which to manage resources."
  type        = string
  default     = "us-west1"
}

variable "zone" {
  description = "The zone in which the compute instance will be created."
  type        = string
  default     = "us-west1-b"
}

variable "image_name" {
  description = "The image from which to initialize this disk."
  type        = string
}

variable "name" {
  description = "The name of the compute instance."
  type        = string
  default     = "nios-gcp-instance"
}

variable "nios_model" {
  description = "The NIOS appliance model used to determine the machine type."
  type        = string
  default     = "IB-V1425"
}

variable "machine_type" {
  description = "The machine type to use for the instance. Used if nios_model is not mapped."
  type        = string
  default     = "n2-standard-4"
}

variable "mgmt_subnet_name" {
  description = "The name of the subnetwork to attach to the primary network interface (nic0)."
  type        = string
}

variable "lan1_subnet_name" {
  description = "The name of the subnetwork to attach to the secondary network interface (nic1)."
  type        = string
}

variable "boot_disk_type" {
  description = "The type of the boot disk."
  type        = string
  default     = "pd-standard"
}

variable "boot_disk_size" {
  description = "The size of the boot disk in GB."
  type        = number
  default     = 250
}

variable "nios_license" {
  description = "The NIOS license string applied during instance initialization."
  type        = string
  default     = "nios IB-V1425 enterprise dns dhcp cloud"
}

variable "remote_console_enabled" {
  description = "Whether to enable remote console access."
  type        = bool
  default     = true
}

variable "default_admin_password" {
  description = "The default admin password for the NIOS instance."
  type        = string
  sensitive   = true
}

variable "service_account_email" {
  description = "The service account e-mail address."
  type        = string
  default     = null
}

variable "service_account_scopes" {
  description = "A list of service scopes to assign to the service account."
  type        = list(string)
  default     = ["https://www.googleapis.com/auth/cloud-platform"]
}

variable "labels" {
  description = "A map of key/value labels to assign to the instance."
  type        = map(string)
  default = {
    product       = "nios"
    dontstop      = "no"
    dontterminate = "yes"
  }
}