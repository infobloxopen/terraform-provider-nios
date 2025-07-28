// Create Network View with Basic Fields
resource "nios_ipam_network_view" "create_network_view" {
  name = "example_network_view"
}

// Create Network View with additional fields
resource "nios_ipam_network_view" "create_network_view_with_additional_fields" {
  name    = "example-network-view2"
  comment = "Example Network View with additional fields"

  remote_reverse_zones = [
    {
      fqdn           = "0.168.192.in-addr.arpa"
      key_type       = "NONE"
      server_address = "192.168.12.12"
    },
    {
      fqdn           = "1.168.192.in-addr.arpa"
      key_type       = "TSIG"
      server_address = "192.168.12.12"
      tsig_key_name  = "aeiou"
      tsig_key_alg   = "HMAC-SHA256"
      tsig_key       = "dGhpc2lzdGVzdHRzaWdrZXk="
    }
  ]

  remote_forward_zones = [
    {
      fqdn           = "fwdzone1.com"
      key_type       = "NONE"
      server_address = "192.168.12.12"
    },
    {
      fqdn           = "fwdzone2.com"
      key_type       = "TSIG"
      server_address = "192.168.12.12"
      tsig_key_name  = "aeiou"
      tsig_key_alg   = "HMAC-SHA256"
      tsig_key       = "dGhpc2lzdGVzdHRzaWdrZXk="
    }
  ]
  mgm_private = true

  extattrs = {
    Site = "location-1"
  }
}
