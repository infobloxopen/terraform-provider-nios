// Create Record NS with Basic Fields
resource "nios_dns_record_ns" "record1" {
  name       = "example.com"
  nameserver = "nsrec1.example.com"
  addresses = [{
    address         = "192.168.1.10"
    auto_create_ptr = false
  }]
  view = "default"
}

// Create Record NS with auto_create_ptr set to true
resource "nios_dns_record_ns" "record2" {
  name       = "example.com"
  nameserver = "nsrec2.example.com"
  addresses = [{
    address         = "192.168.1.11"
    auto_create_ptr = true
  }]
  view = "default"
}
