//Create an NS group delegation 
resource "nios_dns_nsgroup_delegation" "nsgroup_delegation_basic_fields" {
  name = "example_ns_group_del"
  delegate_to = [
    {
      address = "2.3.3.4"
      name    = "delegate_to_ns_group"
    }
  ]
}

//Create an NS group with additional attributes
resource "nios_dns_nsgroup_delegation" "nsgroup_delegation_with_additional_fields" {
  name = "example_ns_group_delegation"
  delegate_to = [
    {
      address = "2.3.3.5"
      name    = "delegate_to_ns_group"
    }
  ]
  comment = "Create NS Group Delegation"
  extattrs = {
    Site = "location-2"
  }
}
