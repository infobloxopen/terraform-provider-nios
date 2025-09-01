// Create NS group forwardingmember with Basic Fields
resource "nios_dns_nsgroup_forwardingmember" "nsgroup_forwardingmember_basic_fields" {
  name = "example_nsgroup_forwardingmember"
  forwarding_servers = [
    {
      name = "infoblox.localdomain"
    }
  ]
  extattrs = {
    Site = "location-1"
  }
}

// Create NS Group forwardingmember with Additional Fields
resource "nios_dns_nsgroup_forwardingmember" "nsgroup_forwardingmember_additional_fields" {
  name    = "example_nsgroup_forwardingmember1"
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
