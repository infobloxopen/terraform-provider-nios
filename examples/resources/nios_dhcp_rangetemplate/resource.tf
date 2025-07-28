// Create DHCP RangeTemplate with required Fields
resource "nios_dhcp_rangetemplate" "range_required_fields" {
  name       = "example_range_template"
  number_of_addresses = 10
  offset = 20
}

// Create DHCP RangeTemplate with additional Fields
resource "nios_dhcp_rangetemplate" "range_additional_fields" {
  name       = "example_range_template_additional_fields"
  number_of_addresses = 10
  offset = 20
  bootfile = "bootfile.iso"
  bootserver = "boot_server1"
  comment = "Example comment for range template"
  email_list = ["abc@example.com", "xyz@example.com"]
  failover_association = "f1"
  server_association_type = "FAILOVER"
  high_water_mark = 100
  high_water_mark_reset = 20
  low_water_mark = 10
  low_water_mark_reset = 5
  nextserver = "next_server1"
  use_nextserver = true
  use_options = true
  use_email_list = true
  use_bootserver = true
  use_bootfile = true
  options = [
    {
      name = "domain-name-servers"
      value = "11.22.1.2"
    },
    {
      name = "time-offset"
      value = "1000"
    },
    {
      name = "domain-name"
      value = "aa.bb.com"
    },
    {
      name = "dhcp-lease-time"
      value = "2000"
    }
  ]
  extattrs = {
    Site = "location-1"
  }
}