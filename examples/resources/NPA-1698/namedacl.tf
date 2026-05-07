// Create Named Access Control Lists (ACLs) with Basic Fields
resource "nios_acl_namedacl" "namedacl_with_basic_fields" {
  name    = var.name
  comment = var.comment

  extattrs = var.extattrs
}


// Create Named Access Control Lists (ACLs) with Additional Fields
resource "nios_acl_namedacl" "namedacl_with_additional_config" {
  name    = var.name2
  comment = var.comment2

  // ACL Entries
  access_list = var.access_list

  extattrs = var.extattrs2
  #exploded_access_list = var.exploded_access_list
}


# data "nios_acl_namedacl" "example_nameacl_data" {
#  
#  //extattrfilters =  {Site = "location-2"}
#  filters = {
#     "name~" = "base.*"
#  }
#  max_results = 2
#  paging      = 1
# 
# }

# output "query_record" {
#   value = data.nios_acl_namedacl.example_nameacl_data.result
# }


# import {
#   to = nios_acl_namedacl.namedacl_with_basic_fields_import
#   id = "namedacl/b25lLmRlZmluZWRfYWNsJDAuc2hyaWthbnQ:shrikant"
# }