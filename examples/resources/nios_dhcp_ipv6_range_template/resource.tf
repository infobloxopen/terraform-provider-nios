// Create an IPV6 DHCP Range Template with required Fields
resource "nios_dhcp_ipv6_range_template" "ipv6_range_template_required_fields" {
  name                = "example_range_template"
  number_of_addresses = 10
  offset              = 20
  // If Terraform Internal ID extensible attribute has cloud access, add `cloud_api_compatible = true`. Otherwise it would throw the below error:
  // Cloud-incompatible template object example_range_template references extensible attribute Terraform Internal ID that is cloud-compatible.
  cloud_api_compatible = false
}

// Create an IPV6 DHCP Range Template with additional Fields
resource "nios_dhcp_ipv6_range_template" "ipv6_range_template_additional_fields" {
  name                = "example_range_template_additional_fields"
  number_of_addresses = 10
  offset              = 20
  // add `cloud_api_compatible = true` if Terraform Internal ID extensible attribute has cloud access
  cloud_api_compatible    = true
  comment                 = "Example comment for ipv6 range template"
  failover_association    = "failover_association"
  server_association_type = "FAILOVER"
  exclude = [
    {
      "number_of_addresses" = 10
      "offset"              = 2
      "comment"             = "Example comment for range template exclude"
    }
  ]
  member = {
    ipv6addr = "2403:8600:80cf:e10c:3a00::1192"
    name     = "infoblox.localdomain"
  }
  delegated_member = []
  use_logic_filter_rules = true
  logic_filter_rules = [
    {
      "filter" = "option_filter"
      "type"   = "Option"
    }
  ]
  option_filter_rules = [
    {
      "filter"     = "option_filter"
      "permission" = "Deny"
    }
  ]
  recycle_leases = false
  use_recycle_leases = true
}
