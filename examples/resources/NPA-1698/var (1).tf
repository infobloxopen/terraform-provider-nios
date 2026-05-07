# Variables for zone_auth

# Required Fields
variable "fqdn" {
  description = "The name of this DNS zone. For a reverse zone, this is in address/cidr format. For other zones, this is in FQDN format."
  type        = string
}

# Optional Fields
variable "view" {
  description = "The name of the DNS view in which the zone resides"
  type        = string
  default     = null
}

variable "zone_format" {
  description = "The format of the zone. Valid values: FORWARD, IPV4, IPV6"
  type        = string
  default     = null
}

variable "comment" {
  description = "Comment for the zone"
  type        = string
  default     = null
}

variable "disable" {
  description = "Determines whether a zone is disabled or not"
  type        = bool
  default     = null
}

variable "locked" {
  description = "If you enable this option, other administrators cannot make conflicting changes"
  type        = bool
  default     = null
}

variable "ns_group" {
  description = "The name server group that serves DNS for this zone"
  type        = string
  default     = null
}

variable "grid_primary" {
  description = "The list of grid primary servers for this zone"
  type        = list(object({
      name = optional(string)
      stealth = optional(bool)
    }))
  default     = null
}

variable "grid_secondaries" {
  description = "The list of grid secondary servers for this zone"
  type        = list(object({
      enable_preferred_primaries = optional(bool)
      grid_replicate = optional(bool)
      lead = optional(bool)
      name = optional(string)
      preferred_primaries = optional(list(object({
        address = optional(string)
        name = optional(string)
        stealth = optional(bool)
        tsig_key = optional(string)
        tsig_key_alg = optional(string)
        tsig_key_name = optional(string)
        use_tsig_key_name = optional(bool)
      })))
      stealth = optional(bool)
    }))
  default     = null
}

variable "external_primaries" {
  description = "The list of external primary DNS servers for this zone"
  type        = list(object({
      address = optional(string)
      name = optional(string)
      stealth = optional(bool)
      tsig_key = optional(string)
      tsig_key_alg = optional(string)
      tsig_key_name = optional(string)
      use_tsig_key_name = optional(bool)
    }))
  default     = null
}

variable "external_secondaries" {
  description = "The list of external secondary DNS servers for this zone"
  type        = list(object({
      address = optional(string)
      name = optional(string)
      stealth = optional(bool)
      tsig_key = optional(string)
      tsig_key_alg = optional(string)
      tsig_key_name = optional(string)
      use_tsig_key_name = optional(bool)
    }))
  default     = null
}

variable "soa_default_ttl" {
  description = "The Time to Live (TTL) value of the SOA record of this zone (in seconds)"
  type        = number
  default     = null
}

variable "soa_expire" {
  description = "The amount of time (in seconds) after which the secondary server stops giving out answers about the zone"
  type        = number
  default     = null
}

variable "soa_negative_ttl" {
  description = "The negative TTL value of the SOA indicates how long a secondary server can cache data for Does Not Respond responses"
  type        = number
  default     = null
}

variable "soa_refresh" {
  description = "The interval (in seconds) at which a secondary server sends a message to the primary server to check that its data is current"
  type        = number
  default     = null
}

variable "soa_retry" {
  description = "How long (in seconds) a secondary server must wait before attempting to recontact the primary server after a connection failure"
  type        = number
  default     = null
}

variable "soa_email" {
  description = "The SOA email value for this zone"
  type        = string
  default     = null
}

variable "soa_serial_number" {
  description = "The serial number in the SOA record"
  type        = number
  default     = null
}

variable "set_soa_serial_number" {
  description = "Flag to set SOA serial number"
  type        = bool
  default     = null
}

variable "allow_active_dir" {
  description = "Active Directory access list"
  type        = list(object({
      address = optional(string)
      permission = optional(string)
    }))
  default     = null
}

variable "allow_fixed_rrset_order" {
  description = "Allow fixed RRset ordering"
  type        = bool
  default     = null
}

variable "allow_gss_tsig_for_underscore_zone" {
  description = "Allow GSS-TSIG for underscore zones"
  type        = bool
  default     = null
}

variable "allow_gss_tsig_zone_updates" {
  description = "Allow GSS-TSIG zone updates"
  type        = bool
  default     = null
}

variable "allow_query" {
  description = "Query access control list"
  type        = list(object({
      address = optional(string)
      permission = optional(string)
      ref = optional(string)
      struct = optional(string)
      tsig_key = optional(string)
      tsig_key_alg = optional(string)
      tsig_key_name = optional(string)
      use_tsig_key_name = optional(bool)
    }))
  default     = null
}

variable "allow_transfer" {
  description = "Zone transfer access control list"
  type        = list(object({
      address = optional(string)
      permission = optional(string)
      ref = optional(string)
      struct = optional(string)
      tsig_key = optional(string)
      tsig_key_alg = optional(string)
      tsig_key_name = optional(string)
      use_tsig_key_name = optional(bool)
    }))
  default     = null
}

variable "allow_update" {
  description = "Dynamic update access control list"
  type        = list(object({
      address = optional(string)
      permission = optional(string)
      ref = optional(string)
      struct = optional(string)
      tsig_key = optional(string)
      tsig_key_alg = optional(string)
      tsig_key_name = optional(string)
      use_tsig_key_name = optional(bool)
    }))
  default     = null
}

variable "allow_update_forwarding" {
  description = "Allow update forwarding"
  type        = bool
  default     = null
}

variable "copy_xfer_to_notify" {
  description = "Copy transfer settings to notify"
  type        = bool
  default     = null
}

variable "create_underscore_zones" {
  description = "Automatically create underscore zones"
  type        = bool
  default     = null
}

variable "ddns_force_creation_timestamp_update" {
  description = "Force DDNS creation timestamp update"
  type        = bool
  default     = null
}

variable "ddns_principal_group" {
  description = "DDNS principal group for tracking"
  type        = string
  default     = null
}

variable "ddns_principal_tracking" {
  description = "Enable DDNS principal tracking"
  type        = bool
  default     = null
}

variable "ddns_restrict_patterns" {
  description = "Enable DDNS restrict patterns"
  type        = bool
  default     = null
}

variable "ddns_restrict_patterns_list" {
  description = "List of DDNS restriction patterns"
  type        = list(string)
  default     = null
}

variable "ddns_restrict_protected" {
  description = "Restrict protected records from DDNS updates"
  type        = bool
  default     = null
}

variable "ddns_restrict_secure" {
  description = "Restrict secure DDNS updates"
  type        = bool
  default     = null
}

variable "ddns_restrict_static" {
  description = "Restrict static records from DDNS updates"
  type        = bool
  default     = null
}

variable "disable_forwarding" {
  description = "Disable zone forwarding"
  type        = bool
  default     = null
}

variable "dns_integrity_enable" {
  description = "Enable DNS integrity checking"
  type        = bool
  default     = null
}

variable "dns_integrity_frequency" {
  description = "DNS integrity check frequency in seconds"
  type        = number
  default     = null
}

variable "dns_integrity_member" {
  description = "Member to perform DNS integrity checks"
  type        = string
  default     = null
}

variable "dns_integrity_verbose_logging" {
  description = "Enable verbose logging for DNS integrity"
  type        = bool
  default     = null
}

variable "dnssec_key_params" {
  description = "DNSSEC key parameters"
  type        = object({
      enable_ksk_auto_rollover = optional(bool)
      ksk_algorithms = optional(list(object({
        algorithm = optional(string)
        size = optional(number)
      })))
      ksk_email_notification_enabled = optional(bool)
      ksk_rollover = optional(number)
      ksk_rollover_notification_config = optional(string)
      ksk_snmp_notification_enabled = optional(bool)
      next_secure_type = optional(string)
      nsec3_iterations = optional(number)
      nsec3_salt_max_length = optional(number)
      nsec3_salt_min_length = optional(number)
      signature_expiration = optional(number)
      zsk_algorithms = optional(list(object({
        algorithm = optional(string)
        size = optional(number)
      })))
      zsk_rollover = optional(number)
      zsk_rollover_mechanism = optional(string)
    })
  default     = null
}

variable "effective_check_names_policy" {
  description = "Check names policy"
  type        = string
  default     = null
}

variable "last_queried_acl" {
  description = "Last queried ACL settings"
  type        = list(object({
      address = optional(string)
      permission = optional(string)
    }))
  default     = null
}

variable "member_soa_mnames" {
  description = "Member SOA MNAME list"
  type        = list(object({
      grid_primary = optional(string)
      mname = optional(string)
      ms_server_primary = optional(string)
    }))
  default     = null
}

variable "ms_ad_integrated" {
  description = "Microsoft Active Directory integrated zone"
  type        = bool
  default     = null
}

variable "ms_allow_transfer" {
  description = "Microsoft DNS allow transfer list"
  type        = list(object({
      address = optional(string)
      permission = optional(string)
    }))
  default     = null
}

variable "ms_allow_transfer_mode" {
  description = "Microsoft DNS allow transfer mode (allowed: ADDRESS_AC, ANY, ANY_NS, NONE)"
  type        = string
  default     = null
}

variable "ms_dc_ns_record_creation" {
  description = "Microsoft DC NS record creation settings"
  type        = list(object({
      address = optional(string)
      comment = optional(string)
    }))
  default     = null
}

variable "ms_ddns_mode" {
  description = "Microsoft DDNS mode"
  type        = string
  default     = null
}

variable "ms_primaries" {
  description = "Microsoft primary DNS servers"
  type        = list(object({
      address = optional(string)
      is_master = optional(bool)
      ns_ip = optional(string)
      ns_name = optional(string)
      stealth = optional(bool)
    }))
  default     = null
}

variable "ms_secondaries" {
  description = "Microsoft secondary DNS servers"
  type        = list(object({
      address = optional(string)
      is_master = optional(bool)
      ns_ip = optional(string)
      ns_name = optional(string)
      stealth = optional(bool)
    }))
  default     = null
}

variable "ms_sync_disabled" {
  description = "Disable Microsoft synchronization"
  type        = bool
  default     = null
}

variable "notify_delay" {
  description = "Notify delay in seconds"
  type        = number
  default     = null
}

variable "prefix" {
  description = "Prefix for the zone"
  type        = string
  default     = null
}

variable "record_name_policy" {
  description = "Record name policy"
  type        = string
  default     = null
}

variable "remove_subzones" {
  description = "Remove subzones when deleting"
  type        = bool
  default     = null
}

variable "restart_if_needed" {
  description = "Restart services if needed"
  type        = bool
  default     = null
}

variable "scavenging_settings" {
  description = "DNS scavenging settings"
  type        = object({
      ea_expression_list = optional(list(object({
        op = optional(string)
        op1 = optional(string)
        op1_type = optional(string)
        op2 = optional(string)
        op2_type = optional(string)
      })))
      enable_auto_reclamation = optional(bool)
      enable_recurrent_scavenging = optional(bool)
      enable_rr_last_queried = optional(bool)
      enable_scavenging = optional(bool)
      enable_zone_last_queried = optional(bool)
      expression_list = optional(list(object({
        op = optional(string)
        op1 = optional(string)
        op1_type = optional(string)
        op2 = optional(string)
        op2_type = optional(string)
      })))
      reclaim_associated_records = optional(bool)
    })
  default     = null
}

variable "srgs" {
  description = "Shared record groups"
  type        = list(string)
  default     = null
}

variable "update_forwarding" {
  description = "Update forwarding settings"
  type        = list(object({
      address = optional(string)
      permission = optional(string)
      struct = optional(string)
      tsig_key = optional(string)
      tsig_key_alg = optional(string)
      tsig_key_name = optional(string)
      use_tsig_key_name = optional(bool)
    }))
  default     = null
}

variable "use_allow_active_dir" {
  description = "Use allow active directory override"
  type        = bool
  default     = null
}

variable "use_allow_query" {
  description = "Use allow query override"
  type        = bool
  default     = null
}

variable "use_allow_transfer" {
  description = "Use allow transfer override"
  type        = bool
  default     = null
}

variable "use_allow_update" {
  description = "Use allow update override"
  type        = bool
  default     = null
}

variable "use_allow_update_forwarding" {
  description = "Use allow update forwarding override"
  type        = bool
  default     = null
}

variable "use_check_names_policy" {
  description = "Use check names policy override"
  type        = bool
  default     = null
}

variable "use_copy_xfer_to_notify" {
  description = "Use copy transfer to notify override"
  type        = bool
  default     = null
}

variable "use_ddns_force_creation_timestamp_update" {
  description = "Use DDNS force creation timestamp update override"
  type        = bool
  default     = null
}

variable "use_ddns_patterns_restriction" {
  description = "Use DDNS patterns restriction override"
  type        = bool
  default     = null
}

variable "use_ddns_principal_security" {
  description = "Use DDNS principal security override"
  type        = bool
  default     = null
}

variable "use_ddns_restrict_protected" {
  description = "Use DDNS restrict protected override"
  type        = bool
  default     = null
}

variable "use_ddns_restrict_static" {
  description = "Use DDNS restrict static override"
  type        = bool
  default     = null
}

variable "use_dnssec_key_params" {
  description = "Use DNSSEC key parameters override"
  type        = bool
  default     = null
}

variable "use_external_primary" {
  description = "Use external primary override"
  type        = bool
  default     = null
}

variable "use_grid_zone_timer" {
  description = "Use grid zone timer override"
  type        = bool
  default     = null
}

variable "use_notify_delay" {
  description = "Use notify delay override"
  type        = bool
  default     = null
}

variable "use_record_name_policy" {
  description = "Use record name policy override"
  type        = bool
  default     = null
}

variable "use_scavenging_settings" {
  description = "Use scavenging settings override"
  type        = bool
  default     = null
}

variable "use_soa_email" {
  description = "Use SOA email override"
  type        = bool
  default     = null
}

variable "extattrs" {
  description = "Extensible attributes"
  type        = map(string)
  default     = null
}
