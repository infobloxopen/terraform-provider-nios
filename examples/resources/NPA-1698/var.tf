variable "name" {
    description = "Name of the Global Smartfolder"
    type        = string
    default     = "example-shared_network"
  
}

variable "name2" {
    description = "Name of the Global Smartfolder"
    type        = string
    default     = "example-shared_network_2"
  
}

variable "networks" {
    description = "List of network references for the shared network"
    type        = list(object({
        ref = string
    }))
    default = null
}

variable "networks2" {
    description = "List of network references for the shared network"
    type        = list(object({
        ref = string
    }))
    default = null
}

variable "network_view" {
    description = "Network view for the shared network"
    type        = string
    default     = null
  
}

variable "extattrs" {
    description = "Extended attributes for the shared network"
    type        = map(string)
    default     = null
  
}

variable "ignore_mac_addresses" {
    description = "List of MAC addresses to ignore"
    type        = list(string)
    default     = null
  
}

variable "use_options" {
    description = "Whether to use DHCP options"
    type        = bool
    default     = null
  
}

variable "options" {
    description = "List of DHCP options"
    type        = list(object({
        name         = string
        num          = number
        value        = string
        vendor_class = string
    }))
    default     = null
  
}

variable "use_logic_filter_rules" {
    description = "Whether to use logic filter rules"
    type        = bool
    default     = null
  
}

variable "logic_filter_rules" {
    description = "List of logic filter rules"
    type        = list(object({
        filter = string
        type   = string
    }))
    default     = null
  
}

variable "comment" {
    description = "Comment for the shared network"
    type        = string
    default     = null
  
}

variable "ddns_server_always_updates" {
    description = "Whether DDNS server always updates"
    type        = bool
    default     = null
  
}

variable "ddns_use_option81" {
    description = "Whether to use option 81 for DDNS"
    type        = bool
    default     = null
  
}

variable "use_ddns_use_option81" {
    description = "Whether to use the ddns_use_option81 setting"
    type        = bool
    default     = null
}

variable "authority" {
    description = "Whether this is an authoritative DHCP server"
    type        = bool
    default     = null
  
}

variable "use_authority" {
    description = "Whether to use the authority setting"
    type        = bool
    default     = null
  
}

variable "ddns_generate_hostname" {
    description = "Whether to generate hostnames for DDNS"
    type        = bool
    default     = null
  
}

variable "use_ddns_generate_hostname" {
    description = "Whether to use the ddns_generate_hostname setting"
    type        = bool
    default     = null
  
}

variable "ddns_ttl" {
    description = "TTL for DDNS records"
    type        = number
    default     = null
  
}

variable "use_ddns_ttl" {
    description = "Whether to use the ddns_ttl setting"
    type        = bool
    default     = null
  
}

variable "ddns_update_fixed_addresses" {
    description = "Whether to update DDNS for fixed addresses"
    type        = bool
    default     = null
  
}

variable "use_ddns_update_fixed_addresses" {
    description = "Whether to use the ddns_update_fixed_addresses setting"
    type        = bool
    default     = null
  
}

variable "deny_bootp" {
    description = "Whether to deny BOOTP requests"
    type        = bool
    default     = null
  
}

variable "use_deny_bootp" {
    description = "Whether to use the deny_bootp setting"
    type        = bool
    default     = null
  
}

variable "disable" {
    description = "Whether to disable the shared network"
    type        = bool
    default     = null
  
}

variable "enable_ddns" {
    description = "Whether to enable DDNS"
    type        = bool
    default     = null
  
}

variable "use_enable_ddns" {
    description = "Whether to use the enable_ddns setting"
    type        = bool
    default     = null
  
}

variable "enable_pxe_lease_time" {
    description = "Whether to enable PXE lease time"
    type        = bool
    default     = null
}

variable "pxe_lease_time" {
    description = "PXE lease time in seconds"
    type        = number
    default     = null
  
}

variable "use_pxe_lease_time" {
    description = "Whether to use the pxe_lease_time setting"
    type        = bool
    default     = null
}

variable "ignore_client_identifier" {
    description = "Whether to ignore client identifiers"
    type        = bool
    default     = null
  
}

variable "use_ignore_client_identifier" {
    description = "Whether to use the ignore_client_identifier setting"
    type        = bool
    default     = null
  
}

variable "ignore_dhcp_option_list_request" {
    description = "Whether to ignore DHCP option list requests"
    type        = bool
    default     = null
  
}

variable "use_ignore_dhcp_option_list_request" {
    description = "Whether to use the ignore_dhcp_option_list_request setting"
    type        = bool
    default     = null
  
}

variable "ignore_id" {
    description = "Client identifier to ignore"
    type        = string
    default     = null
  
}

variable "use_ignore_id" {
    description = "Whether to use the ignore_id setting"
    type        = bool
    default     = null
  
}

variable "lease_scavenge_time" {
    description = "Lease scavenge time in seconds"
    type        = number
    default     = null
  
}

variable "use_lease_scavenge_time" {
    description = "Whether to use the lease_scavenge_time setting"
    type        = bool
    default     = null
}

variable "nextserver" {
    description = "Next server IP address"
    type        = string
    default     = null
  
}

variable "use_nextserver" {
    description = "Whether to use the nextserver setting"
    type        = bool
    default     = null
}

variable "update_dns_on_lease_renewal" {
    description = "Whether to update DNS on lease renewal"
    type        = bool
    default     = null
  
}

variable "use_update_dns_on_lease_renewal" {
    description = "Whether to use the update_dns_on_lease_renewal setting"
    type        = bool
    default     = null
}

variable "bootfile" {
    description = "Bootfile name"
    type        = string
    default     = null
  
}

variable "use_bootfile" {
    description = "Whether to use the bootfile setting"
    type        = bool
    default     = null
}


