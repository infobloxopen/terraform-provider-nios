variable "tenancy_ocid" {
  description = "OCID of the tenancy. Required for creating dynamic groups."
  type        = string
}

variable "compartment_id" {
  description = "OCID of the compartment where NIOS instances reside."
  type        = string
}

variable "dynamic_group_name" {
  description = "Name for the dynamic group."
  type        = string
  default     = "nios-ha-dynamic-group"
}

variable "policy_name" {
  description = "Name for the IAM policy."
  type        = string
  default     = "nios-ha-policy"
}

variable "instance_ocids" {
  description = "List of NIOS HA instance OCIDs to include in the dynamic group. Must contain at least one valid compute instance OCID."
  type        = list(string)

  validation {
    condition     = length(var.instance_ocids) > 0
    error_message = "instance_ocids must contain at least one instance OCID."
  }
}

variable "freeform_tags" {
  description = "Freeform tags to apply to IAM resources."
  type        = map(string)
  default     = {}
}


variable "idcs_endpoint" {
  description = "IDCS endpoint URL for OCI Identity resources (e.g. https://idcs-<region>.oraclecloud.com). Required for creating dynamic groups."
  type        = string
}

variable "identity_domain_name" {
  description = "Name of the Identity Domain (used in policy statements). Usually 'Default' or your custom domain name."
  type        = string
  default     = "Default"
}