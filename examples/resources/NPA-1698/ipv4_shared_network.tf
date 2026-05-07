terraform {
  required_providers {
    nios = {
      source  = "infobloxopen/nios"
      version = "1.1.0"
    }
  }
}

provider "nios" {
  nios_host_url = "https://172.28.82.33"
  nios_username = "admin"
  nios_password = "Infoblox@123"
}

resource "nios_ipam_network" "ipv4_network_parent1" {
  network = "11.0.0.0/24"
}

resource "nios_ipam_network" "ipv4_network_parent2" {
  network = "16.0.0.0/24"
}

resource "nios_ipam_network" "ipv4_network_parent3" {
  network = "17.0.0.0/24"
}



// Create Shared Network with Additional Fields
resource "nios_dhcp_shared_network" "shared_network_additional_fields" {
  name = var.name
  networks = [
    { ref = nios_ipam_network.ipv4_network_parent1.ref },
    { ref = nios_ipam_network.ipv4_network_parent2.ref }
  ]
  extattrs                   = var.extattrs
  ignore_mac_addresses       = var.ignore_mac_addresses
  use_options                = var.use_options
  network_view               = var.network_view
  options                    = var.options
  use_logic_filter_rules     = var.use_logic_filter_rules
  logic_filter_rules         = var.logic_filter_rules
  comment                    = var.comment
  ddns_server_always_updates = var.ddns_server_always_updates
  ddns_use_option81          = var.ddns_use_option81
  use_ddns_use_option81      = var.use_ddns_use_option81

  authority     = var.authority
  use_authority = var.use_authority

  ddns_generate_hostname     = var.ddns_generate_hostname
  use_ddns_generate_hostname = var.use_ddns_generate_hostname

  ddns_ttl     = var.ddns_ttl
  use_ddns_ttl = var.use_ddns_ttl

  ddns_update_fixed_addresses     = var.ddns_update_fixed_addresses
  use_ddns_update_fixed_addresses = var.use_ddns_update_fixed_addresses

  deny_bootp     = var.deny_bootp
  use_deny_bootp = var.use_deny_bootp

  disable = var.disable

  enable_ddns     = var.enable_ddns
  use_enable_ddns = var.use_enable_ddns

  enable_pxe_lease_time = var.enable_pxe_lease_time
  pxe_lease_time        = var.pxe_lease_time
  use_pxe_lease_time    = var.use_pxe_lease_time

  ignore_client_identifier     = var.ignore_client_identifier
  use_ignore_client_identifier = var.use_ignore_client_identifier

  ignore_dhcp_option_list_request     = var.ignore_dhcp_option_list_request
  use_ignore_dhcp_option_list_request = var.use_ignore_dhcp_option_list_request

  ignore_id     = var.ignore_id
  use_ignore_id = var.use_ignore_id

  lease_scavenge_time     = var.lease_scavenge_time
  use_lease_scavenge_time = var.use_lease_scavenge_time

  nextserver     = var.nextserver
  use_nextserver = var.use_nextserver

  update_dns_on_lease_renewal     = var.update_dns_on_lease_renewal
  use_update_dns_on_lease_renewal = var.use_update_dns_on_lease_renewal

  bootfile     = var.bootfile
  use_bootfile = var.use_bootfile

  depends_on = [nios_ipam_network.ipv4_network_parent1, nios_ipam_network.ipv4_network_parent2]


}




# data "nios_dhcp_shared_network" "paged_records" {

#   extattrfilters = {
#     "Site" = "Blr"
#   }

#   max_results = 2
#   paging      = 1
# }


# output "all_records" {
#   value = data.nios_dhcp_shared_network.paged_records.result
# }


resource "nios_dhcp_shared_network" "shared_network_additional_fields1" {
  name = var.name2
  networks = [
    { ref = nios_ipam_network.ipv4_network_parent3.ref }
  ]

  depends_on = [nios_ipam_network.ipv4_network_parent3]
}



# import {
#     to = nios_dhcp_shared_network.shared_network_import
#     id = "sharednetwork/ZG5zLnNoYXJlZF9uZXR3b3JrJHNocmlrYW50MS4w:shrikant1/default"
# }