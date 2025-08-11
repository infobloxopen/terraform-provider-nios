resource "nios_dns_nsgroup_delegation" "create_nsgroup_delegate" {
  name = "example_ns_group_delegation"
  delegate_to = [
    {
      address = "2.3.3.4"
      name    = "delegate_to_ns_group"
  }]
  comment = "Example NS Group Delegation"
  extattrs = {
    Site = "location-1"
  }
}
