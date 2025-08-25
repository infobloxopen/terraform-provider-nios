//Create an NS group delegation 
resource "nios_dns_nsgroup_delegation" "create_nsgroup_delegate" {
  name = "example_ns_group_del"
  delegate_to = [
    {
      address = "2.3.3.4"
      name    = "delegate_to_ns_group"
    }
  ]
}

//Create an NS group with additional attributes
resource "nios_dns_nsgroup_delegation" "create_nsgroup_delegate_with_additional_params" {
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
