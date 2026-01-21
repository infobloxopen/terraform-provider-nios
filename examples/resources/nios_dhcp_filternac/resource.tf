// Create Filter NAC with Basic Fields
resource "nios_dhcp_filternac" "nac_filter_basic_fields" {
  name = "nac_filter_example"
}

// Create Filter NAC with Additional Fields
resource "nios_dhcp_filternac" "nac_filter_additional_fields" {
  name       = "nac_filter_example_2"
  lease_time = 3600
  expression = "(Sophos.ComplianceState=\"NonCompliant\")"
  extattrs = {
    Site = "location-1"
  }
}

// Create another filternac resource with dhcp option dhcp-lease-time time-offset
resource "nios_dhcp_filternac" "nac_filter_with_option" {
  name = "nac_filter_with_option"
  options = [
    {
      name  = "dhcp-lease-time"
      num   = 51
      value = "1200"
    },
    {
      name  = "time-offset"
      num   = 2
      value = "3600"
    }
  ]
}
