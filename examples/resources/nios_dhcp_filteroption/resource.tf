// Create Filteroption with Basic Fields
resource "nios_dhcp_filteroption" "filteroption_basic_fields" {
  name = "filteroption_example"
}

// Create Filteroption with Additional Fields
resource "nios_dhcp_filteroption" "filteroption_additional_fields" {
  name        = "filteroption_example_2"
  lease_time  = 3600
  expression  = "(option domain-name=\"example.com\")"
  option_list = [{ "name" = "time-offset", "num" = 2, "value" = "1200" }]
  next_server = "1.1.1.1"
  bootfile    = "pxelinux.0"
  bootserver  = "1.1.1.2"
  extattrs = {
    Site = "location-1"
  }
}
