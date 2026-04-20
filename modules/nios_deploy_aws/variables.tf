variable "security_group_id" {
  description = "ID of the existing AWS security group."
  type        = string
}

variable "mgmt_subnet_id" {
  description = "ID of the management subnet (ETH0)."
  type        = string
}

variable "lan1_subnet_id" {
  description = "ID of the LAN1 subnet (ETH1)."
  type        = string
}

variable "ami_id" {
  description = "AMI ID for NIOS instance."
  type        = string
  default     = null
}

variable "instance_type" {
  description = "EC2 instance type."
  type        = string
  default     = "r6i.large"
}

variable "key_name" {
  description = "Name of the SSH key pair."
  type        = string
  default     = null
}

variable "availability_zone" {
  description = "AWS availability zone."
  type        = string
  default     = "us-west-1a"
}

variable "volume_size" {
  description = "Size of the root volume in GB."
  type        = number
  default     = 500
}

variable "volume_type" {
  description = "Type of the root volume."
  type        = string
  default     = "gp3"
}

variable "delete_on_termination" {
  description = "Whether to delete the volume on instance termination."
  type        = bool
  default     = true
}

variable "enable_ipv6" {
  description = "Enable IPv6 configuration."
  type        = bool
  default     = false
}

variable "name" {
  description = "Prefix for instance name."
  type        = string
  default     = "nios-aws-instance"
}

variable "tags" {
  description = "Tags to apply to AWS resources."
  type        = map(string)
  default = {
    Name          = "nios-aws-instance"
    dontStop      = "true"
    dontTerminate = "true"
  }
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

variable "ha_enable" {
  description = "Enable HA configuration."
  type        = bool
  default     = false
}

variable "iam_instance_profile" {
  description = "IAM instance profile to attach to the instance."
  type        = string
  default     = null
}

variable "private_ips_count_eth2" {
  description = "Number of IPs to assign to ETH2 (HA interface). Set 1 for secondary IP (VIP) for HA, 0 for no secondary IP."
  type        = number
  default     = 0
}
