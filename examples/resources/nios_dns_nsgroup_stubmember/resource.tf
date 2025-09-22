// Create a NS Group Stub Member with Basic Fields
resource "nios_dns_nsgroup_stubmember" "nsgroup_stubmember_with_basic_fields" {
  name = "stubmember1"
  stub_members = [
    {
      name = "member.com"
    }
  ]
}

// Create a NS Group Stub Member with Additional Fields
resource "nios_dns_nsgroup_stubmember" "nsgroup_stubmember_with_additional_fields" {
  name = "stubmember2"
  stub_members = [
    {
      name = "member.com"
    }
  ]
  comment = "This is a comment"
  extattrs = {
    Site = "location-1"
  }
}
