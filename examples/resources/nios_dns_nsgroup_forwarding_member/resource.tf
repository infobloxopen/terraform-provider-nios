// Create an NS group Forwarding Member with Basic Fields
resource "nios_dns_nsgroup_forwarding_member" "nsgroup_forwarding_member_basic_fields" {
  name = "example_nsgroup_forwarding_member"
  forwarding_servers = [
    {
      name = "infoblox.localdomain"
    }
  ]
  extattrs = {
    Site = "location-1"
  }
}

// Create an NS Group Forwarding Member with Additional Fields
resource "nios_dns_nsgroup_forwarding_member" "nsgroup_forwarding_member_additional_fields" {
  name    = "example_nsgroup_forwarding_member1"
  comment = "nsgroup forwarding member with additional fields"
  forwarding_servers = [
    {
      name            = "infoblox.localdomain"
      forwarders_only = true
      forward_to = [
        {
          name    = "forwarder.localdomain"
          address = "2.3.4.45"
        }
      ]
      use_override_forwarders = true
    }
  ]
  extattrs = {
    Site = "location-1"
  }
}
