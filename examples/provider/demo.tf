# resource "nios_resource_nios_RecordA" "create_record" {
#   name     = "example_test1.example.com"
#   ipv4addr = "10.0.1.2"
#   view     = "default"
#   comment  = "This is a 1 record"
# }
 

# resource "nios_resource_nios_RecordA" "create_record1" {
#   name     = "example_test2.example.com"
#   ipv4addr = "10.1.1.2"
#   view     = "default"
#   comment  = "This is a 2 record"
# }


data "nios_datasource_nios_RecordA" "get_record" {
  extattrfilters = {
    "Site" = "Siteblr"
  }
  filters = {
    "name~" = "example*"
    "view"  = "default"
    "zone"  = "example.com"
    "_max_results" = 100
  }
}

output "data" {
  value = length(data.nios_datasource_nios_RecordA.get_record.result)
}
# resource "nios_resource_nios_RecordA" "create_record2" {
#   name     = "example_test3.example.com"
#   ipv4addr = "1.0.1.2"
#   view     = "default"
# }

# resource "nios_resource_nios_RecordA" "create_record3" {
#   name     = "example_test4.example.com"
#   ipv4addr = "10.0.1.0"
#   view     = "default"
# }

# resource "nios_resource_nios_RecordA" "create_record4" {
#   name     = "example_test5.example.com"
#   ipv4addr = "0.0.1.2"
#   view     = "default"
# }

# resource "nios_resource_nios_RecordA" "create_record5" {
#   name     = "example_test6.example.com"
#   ipv4addr = "2.0.1.0"
#   view     = "default"
# }